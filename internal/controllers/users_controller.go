package controllers

import (
	dto_request "test-api-go/internal/dto/request"
	"test-api-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	UUID        string
	userService *services.UsersService
}

func NewUsersController(us *services.UsersService) UsersController {
	return UsersController{
		UUID:        "1234",
		userService: us,
	}
}

func (uc UsersController) GetUsers(ctx *fiber.Ctx) error {
	response := dto_request.ResponseDTO{
		Data: uc.userService.GetUsers(),
	}

	return ctx.JSON(response)
}

func (uc UsersController) CreateUser(ctx *fiber.Ctx) error {
	body := new(dto_request.CreateUserDTO)

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	result := uc.userService.CreateUser(*body)

	return ctx.SendString(result)
}
