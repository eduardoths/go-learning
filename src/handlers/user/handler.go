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
	interfaces.UserHandler
	router fiber.Router
	userService interfaces.UserService
}

func NewUserHandler(router fiber.Router, servs services.ServiceContainer) interfaces.UserHandler {
	return UserHandler{
		router: router,
		userService: servs.UserService,
	}
}

func (uh UserHandler) SetRoutes(routerGroup string, middleware ...func(*fiber.Ctx) error) {
	var gRouter fiber.Router 
	if middleware != nil {
		gRouter = uh.router.Group(routerGroup, middleware...)
	} else {
		gRouter = uh.router.Group(routerGroup)
	}
	gRouter.Post("/signup", uh.Create)
	gRouter.Post("/login", uh.Login)
}

func (uh UserHandler) Create(ctx *fiber.Ctx) error {
	var user structs.UserRaw
	
	if err := ctx.BodyParser(&user); err != nil {
		log.Printf("Error processing request, Err: %s", err.Error())
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

func (uh UserHandler) Login(ctx *fiber.Ctx) error {
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
