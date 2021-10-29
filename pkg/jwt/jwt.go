package jwt

import (
	"time"

	"github.com/eduardothsantos/go-learning/src/structs"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Data map[string]string
	jwt.RegisteredClaims
}

func GenerateToken(user structs.UserAuthenticated, secret string, expiresAt time.Time) (string, error) {
	dm := make(map[string]string)
	dm["id"] = user.ID
	dm["email"] = user.Email

	claims := CustomClaims{
		dm,
		jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}