package gateway

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
	"github.com/ridwanakf/bridges/twitter/media"
)

type AutobaseGateway struct {
	oauthClient   *http.Client
	twitterClient *twitter.Client
}

func NewAutobaseGateway(cfg config.TwitterKey) *AutobaseGateway {
	// Setup OAuth Client
	consumerKey := cfg.ConsumerKey
	consumerSecret := cfg.ConsumerSecret
	accessToken := cfg.AccessToken
	accessSecret := cfg.AccessTokenSecret

	conf := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	oauthClient := conf.Client(oauth1.NoContext, token)

	return &AutobaseGateway{
		oauthClient:   oauthClient,
		twitterClient: twitter.NewClient(oauthClient),
	}
}

func (g *AutobaseGateway) GetUserInfo() (twitter.User, error) {
	user, _, err := g.twitterClient.Accounts.VerifyCredentials(nil)
	if err != nil {
		return twitter.User{}, err
	}
	return *user, nil
}

func (g *AutobaseGateway) ReadBatchMessage(count int) ([]twitter.DirectMessageEvent, error) {
	messages, _, err := g.twitterClient.DirectMessages.EventsList(
		&twitter.DirectMessageEventsListParams{Count: count},
	)
	if err != nil {
		return []twitter.DirectMessageEvent{}, err
	}
	return messages.Events, nil
}

func (g *AutobaseGateway) ReadMessage(messageID string) (twitter.DirectMessageEvent, error) {
	message, _, err := g.twitterClient.DirectMessages.EventsShow(messageID, nil)
	if err != nil {
		return twitter.DirectMessageEvent{}, err
	}
	return *message, nil
}

func (g *AutobaseGateway) SendMessage(params twitter.DirectMessageEventMessage) error {
	_, _, err := g.twitterClient.DirectMessages.EventsNew(&twitter.DirectMessageEventsNewParams{
		Event: &twitter.DirectMessageEvent{
			Type:    "message_create",
			Message: &params,
		},
	})
	return err
}

func (g *AutobaseGateway) DeleteMessage(messageID string) error {
	_, err := g.twitterClient.DirectMessages.EventsDestroy(messageID)
	return err
}

func (g *AutobaseGateway) Tweet(text string, params *twitter.StatusUpdateParams) (twitter.Tweet, error) {
	tweet, _, err := g.twitterClient.Statuses.Update(text, params)
	if err != nil {
		return twitter.Tweet{}, err
	}
	return *tweet, nil
}

func (g *AutobaseGateway) DownloadMedia(url string, mediaType string) ([]byte, error) {
	var (
		fileByte []byte
		filename string
	)

	switch mediaType {
	case "photo":
		filename = "image.png"
	case "video":
		filename = "video.mp4"
	case "theTypeOfGifs":
		filename = "animation.gif"
	default:
		filename = "image.png"
	}

	file, err := os.Create(filename)
	if err != nil {
		return fileByte, err
	}

	resp, err := g.oauthClient.Get(url)
	if err != nil {
		return fileByte, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fileByte, err
	}
	defer file.Close()

	fileByte, err = ioutil.ReadFile(filename)
	if err != nil {
		return fileByte, err
	}

	err = os.Remove(filename)
	if err != nil {
		log.Println("ERROR WHEN REMOVING FILE")
	}

	return fileByte, nil
}

func (g *AutobaseGateway) UploadMedia(file []byte, mimetype string) (media.Media, error) {
	client := media.NewService(g.oauthClient)
	mediaUpload := &media.UploadParams{
		File:     file,
		MimeType: mimetype,
	}

	response, _, err := client.Upload(mediaUpload)
	if err != nil {
		return media.Media{}, err
	}
	return *response, nil
}
