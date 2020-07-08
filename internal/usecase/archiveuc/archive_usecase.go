package archiveuc

import (
	"errors"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/ridwanakf/autobase-twitter/internal"
	"github.com/ridwanakf/autobase-twitter/internal/entity"
)

type ArchiveUsecase struct {
	repo internal.ArchiveRepo
}

func NewArchiveUsecase(repo internal.ArchiveRepo) *ArchiveUsecase {
	return &ArchiveUsecase{repo: repo}
}

func (uc *ArchiveUsecase) GetAllMessages() ([]entity.Message, error) {
	return uc.repo.GetAllMessages()
}

func (uc *ArchiveUsecase) GetMessageByUserID(userID string) ([]entity.Message, error) {
	if userID == "" {
		return []entity.Message{}, errors.New("userID can not be empty")
	}

	return uc.repo.GetMessageByUserID(userID)
}

func (uc *ArchiveUsecase) GetMessageByUsername(username string) ([]entity.Message, error) {
	if username == "" {
		return []entity.Message{}, errors.New("username can not be empty")
	}

	return uc.repo.GetMessageByUsername(username)
}

func (uc *ArchiveUsecase) SaveMessage(userID string, message entity.Message) error {
	if userID == "" {
		return errors.New("userID can not be empty")
	}

	return uc.repo.SaveMessage(userID, message)
}

func (uc *ArchiveUsecase) ConvertMessage(input twitter.DirectMessageEvent) entity.Message {
	panic("implement me")
}
