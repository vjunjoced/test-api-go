package routes

import (
	"test-api-go/internal/controllers"
	"test-api-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App, userService *services.UsersService) {
	group := app.Group("/users")

	controller := controllers.NewUsersController(userService)

	group.Get("/", controller.GetUsers)
	group.Post("/", controller.CreateUser)
}
