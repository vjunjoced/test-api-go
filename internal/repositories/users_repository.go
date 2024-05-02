package repositories

import dto_request "test-api-go/internal/dto/request"

type UsersRepository struct {
	List []dto_request.CreateUserDTO
}

func NewUsersRepository() *UsersRepository {
	repo := new(UsersRepository)
	repo.List = []dto_request.CreateUserDTO{}

	return repo
}

func (ur *UsersRepository) AddUser(user dto_request.CreateUserDTO) {
	ur.List = append(ur.List, user)
}

func (ur *UsersRepository) GetUsers() []dto_request.CreateUserDTO {
	return ur.List
}
