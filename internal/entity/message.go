package entity

import "time"

type Message struct {
	SenderID         string    `json:"user_id"`
	SenderScreenName string    `json:"username,omitempty"`
	Text             string    `json:"text,omitempty"`
	MediaURL         string    `json:"media_url,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
}
