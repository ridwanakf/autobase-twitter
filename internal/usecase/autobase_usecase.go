package usecase

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/autobase-twitter/internal"
)

type AutobaseUsecase struct {
	gw internal.AutobaseGW
}

func NewAutobaseUsecase(gateway internal.AutobaseGW) *AutobaseUsecase{
	return &AutobaseUsecase{
		gw: gateway,
	}
}

func (u *AutobaseUsecase) GetUserInfo() (twitter.User, error) {
	panic("implement me!")
}

func (u *AutobaseUsecase) ReadBatchMessage(count int) ([]twitter.DirectMessageEvent, error) {
	panic("implement me!")
}

func (u *AutobaseUsecase) ReadMessage(messageID string) (twitter.DirectMessageEvent, error) {
	panic("implement me!")
}

func (u *AutobaseUsecase) GetBatchMessageID(messages []twitter.DirectMessageEvent) ([]string, error) {
	panic("implement me!")
}

func (u *AutobaseUsecase) FilterMessage(keyword string, messages []twitter.DirectMessageEvent) (correctMessages, incorrectMessages []twitter.DirectMessageEvent) {
	panic("implement me!")
}

func (u *AutobaseUsecase) SendMessage(recipientID string, text string, params twitter.DirectMessageEventsNewParams) error {
	panic("implement me!")
}

func (u *AutobaseUsecase) DeleteBatchMessage(messageID []string) error {
	panic("implement me!")
}

func (u *AutobaseUsecase) DeleteMessage(messageID string) error {
	panic("implement me!")
}

func (u *AutobaseUsecase) ProcessBatchTweet(messages []twitter.DirectMessageEvent) ([]twitter.Tweet, error) {
	panic("implement me!")
}

func (u *AutobaseUsecase) ProcessTweet(messages twitter.DirectMessageEvent) (twitter.Tweet, error) {
	panic("implement me!")
}

