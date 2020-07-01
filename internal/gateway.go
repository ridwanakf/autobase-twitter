package internal

import (
	"github.com/dghubble/go-twitter/twitter"
)

type AutobaseGW interface {
	GetUserInfo() (twitter.User, error)
	ReadBatchMessage(count int) ([]twitter.DirectMessageEvent, error)
	ReadMessage(messageID string) (twitter.DirectMessageEvent, error)
	SendMessage(recipientID string, text string, params twitter.DirectMessageEventsNewParams) error
	DeleteMessage(messageID string) error
	Tweet(text string) (twitter.Tweet, error)
	TweetWithMedia(text string, params twitter.StatusUpdateParams) (twitter.Tweet, error)
	DownloadMedia(url string, mediaType string) ([]byte, error)
	UploadMedia(file []byte, mimetype string) twitter.MediaEntity
}
