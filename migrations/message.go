package migrations

import "time"

type Messages struct {
	Base
	Message    string
	SenderID   string
	ReceiverID string
	SentAt     time.Time
}
