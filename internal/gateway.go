package internal

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/bridges/twitter/media"
)

type AutobaseGW interface {
	GetUserInfo() (twitter.User, error)
	ReadBatchMessage(count int) ([]twitter.DirectMessageEvent, error)
	ReadMessage(messageID string) (twitter.DirectMessageEvent, error)
	SendMessage(params twitter.DirectMessageEventMessage) error
	DeleteMessage(messageID string) error
	Tweet(text string) (twitter.Tweet, error)
	TweetWithMedia(text string, params twitter.StatusUpdateParams) (twitter.Tweet, error)
	DownloadMedia(url string, mediaType string) ([]byte, error)
	UploadMedia(file []byte, mimetype string) (media.Media, error)
}
