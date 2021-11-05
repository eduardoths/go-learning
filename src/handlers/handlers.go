package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/eduardothsantos/go-learning/src/handlers/message"
	"github.com/eduardothsantos/go-learning/src/handlers/user"
	"github.com/eduardothsantos/go-learning/src/interfaces"
	"github.com/eduardothsantos/go-learning/src/services"
)

type HandlerContainer struct {
	UserHandler interfaces.UserHandler
	MessageHandler interfaces.MessageHandler
}

func NewHandlerContainer(router fiber.Router, servs services.ServiceContainer) HandlerContainer {
	return HandlerContainer{
		UserHandler: user.NewUserHandler(router, servs),
		MessageHandler: message.NewMessageHandler(router, servs),
	}
}
