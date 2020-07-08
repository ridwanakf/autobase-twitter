package entity

import "time"

type Message struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username,omitempty"`
	Text      string    `json:"text,omitempty"`
	MediaURL  string    `json:"media_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
