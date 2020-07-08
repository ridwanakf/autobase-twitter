package archiveuc

import (
	"errors"
	"strconv"
	"time"

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

func (uc *ArchiveUsecase) SaveMessage(sender entity.User, message twitter.DirectMessageEvent) error {
	if sender.UserID == "" {
		return errors.New("userID can not be empty")
	}

	convertedMessage, err := uc.ConvertMessage(sender, message)
	if err != nil {
		return err
	}

	return uc.repo.SaveMessage(sender, convertedMessage)
}

func (uc *ArchiveUsecase) ConvertMessage(sender entity.User, message twitter.DirectMessageEvent) (entity.Message, error) {
	i, err := strconv.ParseInt(message.CreatedAt, 10, 64)
	if err != nil {
		return entity.Message{}, err
	}
	tm := time.Unix(i, 0)

	return entity.Message{
		SenderID:         sender.UserID,
		SenderScreenName: sender.Username,
		Text:             message.Message.Data.Text,
		MediaURL:         message.Message.Data.Attachment.Media.MediaURL,
		CreatedAt:        tm,
	}, nil
}
