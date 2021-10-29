package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

const saltRounds int = 12

func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), saltRounds)
	return string(hash), err
}

func CompareHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

}
