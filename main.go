package main

import (
	"log"
	"test-api-go/internal/repositories"
	"test-api-go/internal/routes"
	"test-api-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// repository
	usersRepository := repositories.NewUsersRepository()

	// services
	usersService := services.NewUsersService(usersRepository)

	routes.UserRouter(app, usersService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3100"))
}
