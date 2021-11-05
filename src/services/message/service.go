package message

import (
	"github.com/eduardothsantos/go-learning/src/interfaces"
	"github.com/eduardothsantos/go-learning/src/repositories"
	"github.com/eduardothsantos/go-learning/src/structs"
)

type MessageService struct {
	messageRepository interfaces.MessageRepository
}

func NewMessageService(repos repositories.RepositoryContainer) MessageService {
	return MessageService{
		messageRepository: repos.MessageRepository,
	}
}

func (ms MessageService) Create(message structs.Message) error {
	return ms.messageRepository.Create(message)
}

func (ms MessageService) GetAll(user...string) ([]structs.Message, error) {
	return ms.messageRepository.GetAll(user...)
}