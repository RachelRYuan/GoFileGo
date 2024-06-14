package register

import (
	"GOFILEGO/models"
)

// Service interface defines a contract for the service responsible for handling registration-related operations.
type Service interface {
	RegisterService(*RegisterInput) (*models.UserEntity, int)
}

// service struct is the concrete implementation of the Service interface.
type service struct {
	repository Repository
}

// NewRegisterService function is the constructor for the service struct that creates a new instance of the service struct.
func NewRegisterService(repository Repository) *service {
	return &service{repository: repository}
}

// RegisterService method of the service struct implements the RegisterService method from the Service interface.
func (s *service) RegisterService(input *RegisterInput) (*models.UserEntity, int) {
	user := models.UserEntity{
		Email:    input.Email,
		Password: input.Password,
		Username: input.Username,
	}
	return s.repository.RegisterRepository(&user)
}
