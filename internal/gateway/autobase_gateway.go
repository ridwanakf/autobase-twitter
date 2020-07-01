package gateway

import (
	"net/http"

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
	panic("implement me!")
}

func (g *AutobaseGateway) ReadBatchMessage(count int) ([]twitter.DirectMessageEvent, error) {
	panic("implement me!")
}

func (g *AutobaseGateway) ReadMessage(messageID string) (twitter.DirectMessageEvent, error) {
	panic("implement me!")
}

func (g *AutobaseGateway) SendMessage(recipientID string, text string, params twitter.DirectMessageEventsNewParams) error {
	panic("implement me!")
}

func (g *AutobaseGateway) DeleteMessage(messageID string) error {
	panic("implement me!")
}

func (g *AutobaseGateway) Tweet(text string) (twitter.Tweet, error) {
	panic("implement me!")
}

func (g *AutobaseGateway) TweetWithMedia(text string, params twitter.StatusUpdateParams) (twitter.Tweet, error) {
	panic("implement me!")
}

func (g *AutobaseGateway) DownloadMedia(url string, mediaType string) ([]byte, error) {
	panic("implement me!")
}

func (g *AutobaseGateway) UploadMedia(file []byte, mimetype string) media.Media {
	panic("implement me!")
}
