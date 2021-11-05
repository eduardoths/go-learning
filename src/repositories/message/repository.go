package message

import (
	"github.com/eduardothsantos/go-learning/src/structs"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return MessageRepository{
		db: db,
	}
}

func (mr MessageRepository) Create(message structs.Message) error {
	return mr.db.Create(&message).Error
}

func (mr MessageRepository) GetAll(user... string) ([]structs.Message, error) {
	var msg []structs.Message
	query := mr.db.Where("receiver_id IN ? OR sender_id IN ?", user, user)
	results := query.Order("sent_at desc").Find(&msg)
	return msg, results.Error
}