package services

import (
	"github.com/eduardothsantos/go-learning/src/interfaces"
	"github.com/eduardothsantos/go-learning/src/repositories"
	"github.com/eduardothsantos/go-learning/src/services/message"
	"github.com/eduardothsantos/go-learning/src/services/user"
)

type ServiceContainer struct {
	UserService interfaces.UserService
	MessageService interfaces.MessageService
}

func GetServices(repos repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		UserService: user.NewUserService(repos),
		MessageService: message.NewMessageService(repos),
	}
}
