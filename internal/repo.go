package internal

import "github.com/ridwanakf/autobase-twitter/internal/entity"

type AutobaseRepo interface {
	GetAllMessages() ([]entity.Message, error)
	GetMessageByUserID(userID string) ([]entity.Message, error)
	GetMessageByUsername(username string) ([]entity.Message, error)
	SaveMessage(userID string, message entity.Message) error
}
