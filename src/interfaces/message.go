package interfaces

import (
	"github.com/eduardothsantos/go-learning/src/structs"
	"github.com/gofiber/fiber/v2"
)

type MessageHandler interface {
	SetRoutes(routerGroup string, middleware ...func(*fiber.Ctx) error)
}

type MessageService interface {
	Create(message structs.Message) error
	GetAll(user... string) ([]structs.Message, error)
}

type MessageRepository interface {
	Create(message structs.Message) error
	GetAll(user...string) ([]structs.Message, error)
}