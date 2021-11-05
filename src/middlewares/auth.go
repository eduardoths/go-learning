package middlewares

import (
	"net/http"
	"strings"

	"github.com/eduardothsantos/go-learning/internal/config"
	"github.com/eduardothsantos/go-learning/pkg/jwt"
	"github.com/eduardothsantos/go-learning/src/structs"
	"github.com/gofiber/fiber/v2"
)

func EnsureAuth(ctx *fiber.Ctx) error {
	authHeader := string(ctx.Request().Header.Peek("Authorization"))
	authHeaderSplited := strings.Split(authHeader, "Bearer ")
	if len(authHeaderSplited) <= 1 {
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response {
			Data: "Invalid authorization header",
			Tag: "FORBIDDEN",
		})
	}
	token := authHeaderSplited[1]
	user, err := jwt.ParseToken(token, config.GetConfig().JWT_SECRET)
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(structs.Response{
			Data: err.Error(),
			Tag: "FORBIDDEN",
		})
	}
	ctx.Locals("user", user)
	return ctx.Next()
}