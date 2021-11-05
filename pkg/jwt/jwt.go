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

func ParseToken(jwtToken string, secret string) (structs.UserAuthenticated, error) {
	// defer func() {
	// 	log.Println("Panicked parsing token: ", jwtToken)
	// 	return
	// }()
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return structs.UserAuthenticated{}, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return structs.UserAuthenticated{ID: claims.Data["id"], Email: claims.Data["email"] }, nil
	}
	return structs.UserAuthenticated{}, err

}