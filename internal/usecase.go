package internal

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/autobase-twitter/internal/entity"
)

// AutobaseUC is a interface containing a collection of methods for interacting with Direct Messages and Tweets.
type AutobaseUC interface {
	// GetCurrentUserInfo is a method to get current user info
	GetCurrentUserInfo() (twitter.User, error)

	// GetUserInfoByID is a method to get Twitter user info
	GetUserInfoByID(userID string) (twitter.User, error)

	// ReadBatchMessage is a method to get `count` latest Direct Messages
	ReadBatchMessage(count int) ([]twitter.DirectMessageEvent, error)

	// ReadMessage is a method to get a single Direct Message
	ReadMessage(messageID string) (twitter.DirectMessageEvent, error)

	// GetBatchMessageID is a method to get all Direct Messages IDs
	GetBatchMessageID(messages []twitter.DirectMessageEvent) ([]string, error)

	// FilterMessage is a method to filter out Direct Messages that don't contain specific keyword
	FilterMessage(keyword string, messages []twitter.DirectMessageEvent) (correctMessages, incorrectMessages []twitter.DirectMessageEvent)

	// SendBatchMessage is a method to send multiple Direct Message
	SendBatchMessage(params []twitter.DirectMessageEventMessage)

	// SendMessage is a method to send Direct Message
	SendMessage(param twitter.DirectMessageEventMessage) error

	// DeleteBatchMessage is a method to delete multiple Direct Messages
	DeleteBatchMessage(messageIDs []string)

	// DeleteBatchMessage is a method to delete a single Direct Message
	DeleteMessage(messageID string) error

	// ProcessTweet is a method to process a single Tweet
	ProcessTweet(messages twitter.DirectMessageEvent) (twitter.Tweet, error)
}

// ArchiveUC is a interface containing a collection of methods for saving Direct Messages in database for archive.
type ArchiveUC interface {
	GetAllMessages() ([]entity.Message, error)
	GetMessageByUserID(userID string) ([]entity.Message, error)
	GetMessageByUsername(username string) ([]entity.Message, error)
	SaveMessage(sender entity.User, message twitter.DirectMessageEvent) error
	ConvertMessage(sender entity.User, message twitter.DirectMessageEvent) (entity.Message, error)
}

// RulesUC is a interface containing a collection of methods for filtering which account that can use bot's feature
type RulesUC interface {
	ResolveRules(param entity.RulesParam) (bool, error)
	IsFollower(relation twitter.Relationship) bool
	IsFollowing(relation twitter.Relationship) bool
	FollowersCount(userID string) (int32, error)
}
