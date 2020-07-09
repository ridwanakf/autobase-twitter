package worker

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/autobase-twitter/internal/app"
	"github.com/ridwanakf/autobase-twitter/internal/delivery/worker/service"
)

func Start(app *app.AutobaseApp) {
	config := app.Cfg

	// Populate params from config
	messageCount := config.Params.MessageCount
	keyword := config.Params.Keyword

	successMessageResponse := config.Params.MessageResponse.Success
	failedMessageResponse := config.Params.MessageResponse.Failed
	incorrectMessageResponse := config.Params.MessageResponse.Incorrect

	intervalDuration := time.Second * time.Duration(config.Params.DelayDuration.Interval)
	sleepDuration := time.Second * time.Duration(config.Params.DelayDuration.Sleep)
	ratelimitDuration := time.Second * time.Duration(config.Params.DelayDuration.RateLimit)

	useArchive := config.UseArchive

	// Init worker service
	svc := service.GetServices(app)

	// Get user info
	user, err := svc.GetUserInfo()
	if err != nil {
		log.Fatal(err)
	}
	userID := strconv.FormatInt(user.ID, 10)

	// Start loop
	for {
		messages, err := svc.GetMessages(messageCount)
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "88 rate limit exceeded") {
				log.Printf("Rate limit has been reached! Will sleep for %v.", ratelimitDuration)
				time.Sleep(ratelimitDuration)
				continue
			} else {
				log.Fatal(err)
			}
		}

		if len(messages) > 0 {
			var (
				incorrectMessages []twitter.DirectMessageEvent
				successMessages   []twitter.DirectMessageEvent
				failedMessages    []twitter.DirectMessageEvent
			)

			// Filter correct formatted messages
			correctMessages, temp := svc.FilterMessage(keyword, messages)
			incorrectMessages = temp

			for _, message := range correctMessages {
				// Delete our own message
				id := message.ID
				if userID == message.Message.SenderID {
					err = svc.DeleteMessage(id)
					if err != nil {
						log.Println(err)
					}
					continue
				}

				// Send tweet
				_, err = svc.SendTweet(message)
				if err != nil {
					log.Printf("error when tweeting %+v", err)
					failedMessages = append(failedMessages, message)
					continue
				}
				successMessages = append(successMessages, message)
			}
			// Delete, send message response, and save message (if useArchive is true) in goroutine
			go func() {
				svc.CleanMessages(successMessages, successMessageResponse, useArchive)
				svc.CleanMessages(failedMessages, failedMessageResponse, useArchive)
				svc.CleanMessages(incorrectMessages, incorrectMessageResponse, useArchive)
			}()
		} else {
			log.Printf("There is no new DM. Will sleep for %v.", sleepDuration)
			time.Sleep(sleepDuration)
			continue
		}
		log.Printf("Success retrieving %d DMs. Will pause for %v.", int(math.Min(float64(len(messages)), float64(messageCount))), intervalDuration)
		time.Sleep(intervalDuration)
	}

}
