package user

import (
	"time"

	"github.com/eduardothsantos/go-learning/internal/bcrypt"
	"github.com/eduardothsantos/go-learning/internal/config"
	"github.com/eduardothsantos/go-learning/pkg/jwt"
	"github.com/eduardothsantos/go-learning/src/interfaces"
	"github.com/eduardothsantos/go-learning/src/repositories"
	"github.com/eduardothsantos/go-learning/src/structs"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(repos repositories.RepositoryContainer) UserService {
	return UserService{
		userRepository: repos.UserRepository,
	}
}

func (us UserService) Create(user structs.UserRaw) error {
	var treatedUser structs.User
	passwordHash, err := bcrypt.Hash(user.Password)
	if err != nil {
		return err
	}
	treatedUser = structs.User{
		Name: user.Name,
		Email: user.Email,
		PasswordHash: passwordHash,
	}
	return us.userRepository.Create(treatedUser)
}

func (us UserService) Login(user structs.UserLogin) (*string, error) {
	userPreAuth, err := us.userRepository.Login(user.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHash(user.Password, userPreAuth.PasswordHash)
	if err != nil {
		return nil, err
	}
	userAuth := structs.UserAuthenticated{
		ID: userPreAuth.ID,
		Email: userPreAuth.Email,
	}
	token, err := jwt.GenerateToken(userAuth, config.GetConfig().JWT_SECRET, time.Now().Add(8 * time.Hour))
	return &token, err
}