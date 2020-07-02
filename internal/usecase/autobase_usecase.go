package usecase

import (
	"errors"
	"log"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/autobase-twitter/internal"
)

type AutobaseUsecase struct {
	gw internal.AutobaseGW
}

func NewAutobaseUsecase(gateway internal.AutobaseGW) *AutobaseUsecase {
	return &AutobaseUsecase{
		gw: gateway,
	}
}

func (u *AutobaseUsecase) GetUserInfo() (twitter.User, error) {
	return u.gw.GetUserInfo()
}

func (u *AutobaseUsecase) ReadBatchMessage(count int) ([]twitter.DirectMessageEvent, error) {
	return u.gw.ReadBatchMessage(count)
}

func (u *AutobaseUsecase) ReadMessage(messageID string) (twitter.DirectMessageEvent, error) {
	return u.gw.ReadMessage(messageID)
}

func (u *AutobaseUsecase) GetBatchMessageID(messages []twitter.DirectMessageEvent) ([]string, error) {
	var ids []string

	if len(messages) <= 0 {
		return ids, errors.New("message is empty")
	}

	for _, message := range messages {
		ids = append(ids, message.ID)
	}
	return ids, nil
}

func (u *AutobaseUsecase) FilterMessage(keyword string, messages []twitter.DirectMessageEvent) (correctMessages, incorrectMessages []twitter.DirectMessageEvent) {
	var (
		correct   []twitter.DirectMessageEvent
		incorrect []twitter.DirectMessageEvent
	)

	for _, message := range messages {
		if strings.Contains(message.Message.Data.Text, keyword) {
			correct = append(correct, message)
		} else {
			incorrect = append(incorrect, message)
		}
	}
	return correct, incorrect
}

func (u *AutobaseUsecase) SendBatchMessage(params []twitter.DirectMessageEventMessage) {
	for _, param := range params {
		err := u.SendMessage(param)
		if err != nil {
			log.Printf("error when sending message: %+v", err)
		}
	}
}

func (u *AutobaseUsecase) SendMessage(param twitter.DirectMessageEventMessage) error {
	return u.gw.SendMessage(param)
}

func (u *AutobaseUsecase) DeleteBatchMessage(messageIDs []string) {
	for _, messageID := range messageIDs {
		err := u.DeleteMessage(messageID)
		if err != nil {
			log.Printf("error when deleting message: %+v", err)
		}
	}
}

func (u *AutobaseUsecase) DeleteMessage(messageID string) error {
	return u.gw.DeleteMessage(messageID)
}

func (u *AutobaseUsecase) ProcessTweet(message twitter.DirectMessageEvent) (twitter.Tweet, error) {
	var (
		tweetParams *twitter.StatusUpdateParams
		status      string
	)

	if message.Message.Data.Attachment != nil {
		tweetParams = &twitter.StatusUpdateParams{}
		mediaURL := message.Message.Data.Attachment.Media.MediaURLHttps
		mediaType := message.Message.Data.Attachment.Type

		file, err := u.gw.DownloadMedia(mediaURL, mediaType)
		if err != nil {
			return twitter.Tweet{}, err
		}

		response, err := u.gw.UploadMedia(file, mediaURL)
		if err != nil {
			log.Fatal(err)
		}

		tweetParams.MediaIds = []int64{response.MediaID}
		status = strings.ReplaceAll(message.Message.Data.Text, message.Message.Data.Attachment.Media.URL, "")
	} else {
		status = message.Message.Data.Text
	}

	tweet, err := u.gw.Tweet(status, tweetParams)
	if err != nil {
		return twitter.Tweet{}, err
	}
	return tweet, nil
}
