package service

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/app"
	"github.com/ridwanakf/autobase-twitter/internal/entity"
)

type WorkerService struct {
	autobaseUC internal.AutobaseUC
	archiveUC  internal.ArchiveUC
}

func NewWorkerService(app *app.AutobaseApp) *WorkerService {
	return &WorkerService{
		autobaseUC: app.UseCases.AutobaseUC,
		archiveUC:  app.UseCases.ArchiveUC,
	}
}

func (s *WorkerService) GetUserInfo() (twitter.User, error) {
	return s.autobaseUC.GetCurrentUserInfo()
}

func (s *WorkerService) GetMessages(count int) ([]twitter.DirectMessageEvent, error) {
	return s.autobaseUC.ReadBatchMessage(count)
}

func (s *WorkerService) GetBatchMessageID(messages []twitter.DirectMessageEvent) ([]string, error) {
	return s.autobaseUC.GetBatchMessageID(messages)
}

func (s *WorkerService) FilterMessage(keyword string, messages []twitter.DirectMessageEvent) (correctMessages, incorrectMessages []twitter.DirectMessageEvent) {
	return s.autobaseUC.FilterMessage(keyword, messages)
}

func (s *WorkerService) SendBatchMessage(params []twitter.DirectMessageEventMessage) {
	s.autobaseUC.SendBatchMessage(params)
}

func (s *WorkerService) DeleteBatchMessage(messageIDs []string) {
	s.autobaseUC.DeleteBatchMessage(messageIDs)
}

func (s *WorkerService) DeleteMessage(messageID string) error {
	return s.autobaseUC.DeleteMessage(messageID)
}

func (s *WorkerService) SendTweet(message twitter.DirectMessageEvent) (twitter.Tweet, error) {
	return s.autobaseUC.ProcessTweet(message)
}

func (s *WorkerService) CleanMessages(messages []twitter.DirectMessageEvent, response string, archive bool) {
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
	if archive {
		s.SaveBatchMessage(messages)
	}
}

func (s *WorkerService) SaveBatchMessage(messages []twitter.DirectMessageEvent) {
	for _, message := range messages {
		err := s.SaveMessage(message)
		if err != nil {
			log.Printf("error when saving message: %+v", err)
		}
	}
}

func (s *WorkerService) SaveMessage(message twitter.DirectMessageEvent) error {
	u, err := s.autobaseUC.GetUserInfoByID(message.Message.SenderID)
	if err != nil {
		return err
	}

	user := entity.User{
		UserID:   u.IDStr,
		Username: u.ScreenName,
	}

	err = s.archiveUC.SaveMessage(user, message)
	if err != nil {
		return err
	}

	return nil
}
