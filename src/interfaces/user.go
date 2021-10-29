package interfaces

import (
	"github.com/eduardothsantos/go-learning/src/structs"
)

type UserService interface {
	Create(user structs.UserRaw) error
	Login(user structs.UserLogin) (*string, error)
	// ListAllUsers() ([]structs.User, error)
}

type UserRepository interface {
	Create(user structs.User) error
	Login(email string) (structs.UserPreAuth, error) 
	// ListAllUsers(user structs.User) ([]structs.User, error)
}
