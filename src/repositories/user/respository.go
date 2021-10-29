package user

import (
	"github.com/eduardothsantos/go-learning/src/structs"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type User struct {
	ID string
	Email string
	PasswordHash string
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (ur UserRepository) Create(user structs.User) error {
	return ur.db.Create(user).Error
}

func (ur UserRepository) Login(email string) (structs.UserPreAuth, error) {
	user := User{}
	result := ur.db.First(&user, "email = ?", email)
	userPreAuth := structs.UserPreAuth{
		ID: user.ID,
		Email: user.Email,
		PasswordHash: user.PasswordHash,
	}
	return userPreAuth, result.Error
}