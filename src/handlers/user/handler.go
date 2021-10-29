package user

import (
	"log"
	"net/http"

	"github.com/eduardothsantos/go-learning/src/interfaces"
	"github.com/eduardothsantos/go-learning/src/services"
	"github.com/eduardothsantos/go-learning/src/structs"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	router fiber.Router
	userService interfaces.UserService
}

func NewUserHandler(router fiber.Router, servs services.ServiceContainer) UserHandler {
	return UserHandler{
		router: router,
		userService: servs.UserService,
	}
}

func (uh UserHandler) SetRoutes() {
	gRouter := uh.router.Group("/users")
	gRouter.Post("/signup", uh.CreateUser)
	gRouter.Post("/login", uh.LoginUser)
}

func (uh UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var user structs.UserRaw
	
	if err := ctx.BodyParser(&user); err != nil {
		log.Fatalf("Error processing request, Err: %s", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response{
			Data: err.Error(),
			Tag: "STATUS_BAD_REQUEST",
		})
	}

	if err := uh.userService.Create(user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response{
			Data: err.Error(),
			Tag: "INTERNAL_SERVER_ERROR",
		})
	}

	return ctx.Status(http.StatusCreated).JSON(structs.Response {
		Data: "User created",
	})
}

func (uh UserHandler) LoginUser(ctx *fiber.Ctx) error {
	var user structs.UserLogin;

	if err := ctx.BodyParser(&user); err != nil {
		log.Fatalf("Error processing request, Err: %s", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response{
			Data: err.Error(),
			Tag: "STATUS_BAD_REQUEST",
		})
	}

	token, err := uh.userService.Login(user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response {
			Data: err.Error(),
			Tag: "INTERNAL_SERVER_ERROR",
		})
	}
	return ctx.Status(http.StatusAccepted).JSON(structs.Response {
		Data: *token,
	})
}
