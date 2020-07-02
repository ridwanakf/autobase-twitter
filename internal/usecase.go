package internal

import (
	"github.com/dghubble/go-twitter/twitter"
)

type AutobaseUC interface {
	// GetUserInfo is a method to get current user info
	GetUserInfo() (twitter.User, error)

	// ReadBatchMessage is a method to get `count` latest Direct Messages
	ReadBatchMessage(count int) ([]twitter.DirectMessageEvent, error)

	// ReadMessage is a method to get a single Direct Message
	ReadMessage(messageID string) (twitter.DirectMessageEvent, error)

	// GetBatchMessageID is a method to get all Direct Messages IDs
	GetBatchMessageID(messages []twitter.DirectMessageEvent) ([]string, error)

	// FilterMessage is a method to filter out Direct Messages that don't contain specific keyword
	FilterMessage(keyword string, messages []twitter.DirectMessageEvent) (correctMessages, incorrectMessages []twitter.DirectMessageEvent)

	// SendMessage is a method to send Direct Message
	SendMessage(recipientID string, text string, params twitter.DirectMessageEventMessage) error

	// DeleteBatchMessage is a method to delete multiple Direct Messages
	DeleteBatchMessage(messageIDs []string)

	// DeleteBatchMessage is a method to delete a single Direct Message
	DeleteMessage(messageID string) error

	// ProcessTweet is a method to process a single Tweet
	ProcessTweet(messages twitter.DirectMessageEvent) (twitter.Tweet, error)
}
