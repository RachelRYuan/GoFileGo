package loginAuth

import (
	"GOFILEGO/models"
)

// Service interface defines a contract for the service
// it has only one function LoginService which takes the LoginInput and returns the UserEntity and int
type Service interface {
	LoginService(input *LoginInput) (*models.UserEntity, int)
}

// service struct is the concrete implementation of the Service interface.
type service struct {
	repository Repository
}

// NewServiceLogin function is the constructor for the service struct that
// creates a new instance of the service struct
func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

// LoginService method of the service struct implements the LoginService method
// from the Service interface.
func (s *service) LoginService(input *LoginInput) (*models.UserEntity, int) {
	userEntity := models.UserEntity{
		Email:    input.Email,
		Password: input.Password,
	}
	return s.repository.LoginRepository(&userEntity)
}
