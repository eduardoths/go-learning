package interfaces

import (
	"github.com/eduardothsantos/go-learning/src/structs"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	SetRoutes(routerGroup string, middleware ...func(*fiber.Ctx) error)
}
type UserService interface {
	Create(user structs.UserRaw) error
	Login(user structs.UserLogin) (*string, error)
}

type UserRepository interface {
	Create(user structs.User) error
	Login(email string) (structs.UserPreAuth, error) 
}
