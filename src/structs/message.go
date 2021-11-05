package structs

import "time"

type Message struct {
	SenderID   string
	ReceiverID string    `json:"receiver_id"`
	Message    string    `json:"message"`
	SentAt     time.Time
}
