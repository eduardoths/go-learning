package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eduardothsantos/go-learning/internal/config"
	"github.com/eduardothsantos/go-learning/src/handlers"
	"github.com/eduardothsantos/go-learning/src/middlewares"
	"github.com/eduardothsantos/go-learning/src/repositories"
	"github.com/eduardothsantos/go-learning/src/services"
	"github.com/eduardothsantos/go-learning/src/structs"

	"github.com/gofiber/fiber/v2"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	config := config.GetConfig()
	db := config.GetDBConnector()
	app := fiber.New()

	repositoriesContainer := repositories.GetRepositories(db)
	servicesContainer := services.GetServices(repositoriesContainer)
	handlersContainer := handlers.NewHandlerContainer(app, servicesContainer)
	middlewaresContainer := middlewares.NewMiddlewareContainer()
	handlersContainer.UserHandler.SetRoutes("/users")
	handlersContainer.MessageHandler.SetRoutes("/messages", middlewaresContainer.AuthMiddleware.Authenticate)

	app.Get("/health", health)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error starting server, Err= %s", err.Error())
	}

}

func health(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(structs.Response{
		Data: "Ok",
	})
}
