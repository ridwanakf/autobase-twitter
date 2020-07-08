package service

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app"
)

type WorkerService struct {
	uc internal.AutobaseUC
}

func NewWorkerService(app *app.AutobaseApp) *WorkerService {
	return &WorkerService{
		uc: app.UseCases.AutobaseUC,
	}
}

func (s *WorkerService) GetUserInfo() (twitter.User, error) {
	return s.uc.GetCurrentUserInfo()
}

func (s *WorkerService) GetMessages(count int) ([]twitter.DirectMessageEvent, error) {
	return s.uc.ReadBatchMessage(count)
}

func (s *WorkerService) GetBatchMessageID(messages []twitter.DirectMessageEvent) ([]string, error) {
	return s.uc.GetBatchMessageID(messages)
}

func (s *WorkerService) FilterMessage(keyword string, messages []twitter.DirectMessageEvent) (correctMessages, incorrectMessages []twitter.DirectMessageEvent) {
	return s.uc.FilterMessage(keyword, messages)
}

func (s *WorkerService) SendBatchMessage(params []twitter.DirectMessageEventMessage) {
	s.uc.SendBatchMessage(params)
}

func (s *WorkerService) DeleteBatchMessage(messageIDs []string) {
	s.uc.DeleteBatchMessage(messageIDs)
}

func (s *WorkerService) DeleteMessage(messageID string) error {
	return s.uc.DeleteMessage(messageID)
}

func (s *WorkerService) SendTweet(message twitter.DirectMessageEvent) (twitter.Tweet, error) {
	return s.uc.ProcessTweet(message)
}

func (s *WorkerService) CleanMessages(messages []twitter.DirectMessageEvent, response string) {
	var messageResponse []twitter.DirectMessageEventMessage
	messageIDs, _ := s.GetBatchMessageID(messages)

	for _, message := range messages {
		messageResponse = append(messageResponse, twitter.DirectMessageEventMessage{
			Target: &twitter.DirectMessageTarget{
				RecipientID: message.Message.SenderID,
			},
			Data: &twitter.DirectMessageData{
				Text: response,
			},
		})
	}
	s.SendBatchMessage(messageResponse)
	s.DeleteBatchMessage(messageIDs)
}