package repositories

import (
	"gorm.io/gorm"

	"github.com/eduardothsantos/go-learning/src/interfaces"
	"github.com/eduardothsantos/go-learning/src/repositories/user"
)

type RepositoryContainer struct {
	UserRepository interfaces.UserRepository
}

func GetRepositories(db *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		UserRepository: user.NewUserRepository(db),
	}
}
