package middlewares

import (
	"github.com/eduardothsantos/go-learning/src/interfaces"
	"github.com/eduardothsantos/go-learning/src/middlewares/auth"
)


type MiddlewareContainer struct {
	AuthMiddleware interfaces.AuthMiddleware
}

func NewMiddlewareContainer() MiddlewareContainer {
	return MiddlewareContainer{
		AuthMiddleware: auth.NewAuthMiddleware(),
	}
}