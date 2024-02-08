package auth

import (
	"awesomeProject/domain"
)

type AuthRepository interface {
	LoginUser(loginForm *domain.LoginFormDTO) (*domain.UserDTO, error)
	CreateUser(user *domain.User, activationLink string) (*domain.UserDTO, error)
	ActivateUser(activationLink string) (*domain.UserDTO, error)
}
