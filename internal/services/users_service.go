package services

import (
	dto_request "test-api-go/internal/dto/request"
	"test-api-go/internal/repositories"
)

type UsersService struct {
	repo *repositories.UsersRepository
}

func NewUsersService(repo *repositories.UsersRepository) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (us UsersService) GetUsers() []dto_request.CreateUserDTO {
	return us.repo.GetUsers()
}

func (us UsersService) CreateUser(body dto_request.CreateUserDTO) string {
	us.repo.AddUser(body)

	return "Create User"
}
