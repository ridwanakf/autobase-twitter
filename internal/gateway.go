package internal

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/bridges/twitter/media"
)

type AutobaseGW interface {
	GetCurrentUserInfo() (twitter.User, error)
	GetUserInfoByID(userID string) (twitter.User, error)
	ReadBatchMessage(count int) ([]twitter.DirectMessageEvent, error)
	ReadMessage(messageID string) (twitter.DirectMessageEvent, error)
	SendMessage(params twitter.DirectMessageEventMessage) error
	DeleteMessage(messageID string) error
	Tweet(text string, params *twitter.StatusUpdateParams) (twitter.Tweet, error)
	DownloadMedia(url string, mediaType string) ([]byte, error)
	UploadMedia(file []byte, mimetype string) (media.Media, error)
	UsersRelationship(sourceID string, targetID string) (twitter.Relationship, error)
}
