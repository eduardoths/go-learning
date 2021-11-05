package message

import (
	"log"
	"net/http"
	"time"

	"github.com/eduardothsantos/go-learning/src/interfaces"
	"github.com/eduardothsantos/go-learning/src/services"
	"github.com/eduardothsantos/go-learning/src/structs"
	"github.com/gofiber/fiber/v2"
)

type MessageHandler struct {
	interfaces.MessageHandler
	router fiber.Router
	MessageService interfaces.MessageService
}

func NewMessageHandler(router fiber.Router, servs services.ServiceContainer) interfaces.MessageHandler {
	return MessageHandler{
		router: router,
		MessageService: servs.MessageService,
	}
}

func (mh MessageHandler) SetRoutes(routerGrup string, middleware ...func(*fiber.Ctx) error) {
	var gRouter fiber.Router
	if middleware != nil {
		gRouter = mh.router.Group(routerGrup, middleware...)
	} else {
		gRouter = mh.router.Group(routerGrup)
	}
	gRouter.Post("/", mh.Create)
	gRouter.Get("/:id", mh.GetAll)
}

func (mh MessageHandler) Create(ctx *fiber.Ctx) error {
	var message structs.Message

	user := ctx.Locals("user")
	defer func () {
		if r := recover(); r != nil {
			log.Printf("Error: type of ctx.Locals(\"users\") != structs.UserAuthenticated")
		}
	}()
	if err := ctx.BodyParser(&message); err != nil {
		log.Printf("Error processing request, Err: %s", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response{
			Data: err.Error(),
			Tag: "STATUS_BAD_REQUEST",
		})
	}
	message.SenderID = user.(structs.UserAuthenticated).ID
	message.SentAt = time.Now()

	if err := mh.MessageService.Create(message); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response {
			Data: err.Error(),
			Tag: "INTERNAL_SERVER_ERROR",
		})
	}
	return ctx.Status(http.StatusAccepted).JSON(structs.Response {
		Data: "Message sent",
	})	
}

func (mh MessageHandler) GetAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	defer func () {
		if r := recover(); r != nil {
			log.Printf("Error: type of ctx.Locals(\"users\") != structs.UserAuthenticated")
		}
	}()

	searchingUser := user.(structs.UserAuthenticated).ID
	searchedUser := ctx.Params("id")
	if searchedUser == "" {
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response {
			Data: "Invalid user",
			Tag: "STATUS_BAD_REQUEST",
		})
	}
	messages, err := mh.MessageService.GetAll(searchingUser, searchedUser)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(structs.Response{
			Data: err.Error(),
			Tag: "INTERNAL_SERVER_ERROR",
		})
	}
	return ctx.Status(http.StatusAccepted).JSON(structs.Response {
		Data: messages,
	})
	
}