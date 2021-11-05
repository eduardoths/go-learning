package interfaces

import "github.com/gofiber/fiber/v2"

type AuthMiddleware interface {
	Authenticate(ctx *fiber.Ctx) error
}