package dto_request

type CreateUserDTO struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type ResponseDTO struct {
	Data interface{} `json:"data"`
}
