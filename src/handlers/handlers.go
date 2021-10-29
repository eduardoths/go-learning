package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/eduardothsantos/go-learning/src/handlers/user"
	"github.com/eduardothsantos/go-learning/src/services"
)

func NewHandlerContainer(router fiber.Router, servs services.ServiceContainer) {
	user.NewUserHandler(router, servs).SetRoutes()
}
