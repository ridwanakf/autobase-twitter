package firebase

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/ridwanakf/autobase-twitter/internal/app/config"
	"github.com/ridwanakf/autobase-twitter/internal/entity"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

type RealtimeDatabase struct {
	ctx    context.Context
	client *db.Client
}

func NewRealtimeDatabase(cfg config.Firebase) (*RealtimeDatabase, error) {
	// Construct credential in json format
	credByte, err := json.Marshal(cfg.FirebaseAdmin)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	opt := option.WithCredentialsJSON(credByte)

	// Check if credential is correct
	_, err = transport.Creds(ctx, opt)
	if err != nil {
		return nil, err
	}

	// Firebase config
	firebaseConfig := &firebase.Config{
		DatabaseURL: cfg.DatabaseURL,
	}

	// Init Firebase Admin
	app, err := firebase.NewApp(ctx, firebaseConfig, opt)
	if err != nil {
		return nil, err
	}

	// Init Firebase RTDB
	rtdb, err := app.Database(ctx)
	if err != nil {
		return nil, err
	}

	return &RealtimeDatabase{
		ctx:    ctx,
		client: rtdb,
	}, nil
}

func (db *RealtimeDatabase) GetAllMessages() ([]entity.Message, error) {
	var messages []entity.Message

	ref := db.client.NewRef(MessageRef)
	if err := ref.Get(db.ctx, &messages); err != nil {
		return messages, err
	}

	return messages, nil
}

func (db *RealtimeDatabase) GetMessageByUserID(userID string) ([]entity.Message, error) {
	var messages []entity.Message

	ref := db.client.NewRef(MessageRef)
	uidRef := ref.Child(userID)
	if err := uidRef.Get(db.ctx, &messages); err != nil {
		return messages, err
	}

	return messages, nil
}

func (db *RealtimeDatabase) GetMessageByUsername(username string) ([]entity.Message, error) {
	var (
		messages []entity.Message
		userID   string
	)

	ref := db.client.NewRef(UsernameMapRef)
	if err := ref.Get(db.ctx, &userID); err != nil {
		return messages, err
	}

	return db.GetMessageByUserID(userID)
}

func (db *RealtimeDatabase) SaveMessage(userID string, message entity.Message) error {
	ref := db.client.NewRef(MessageRef)
	uidRef := ref.Child(userID)
	if err := uidRef.Set(db.ctx, message); err != nil {
		return err
	}
	return nil
}
