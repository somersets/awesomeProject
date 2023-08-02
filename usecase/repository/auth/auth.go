package auth

import (
	"awesomeProject/domain"
)

type AuthRepository interface {
	LoginUser(loginForm *domain.LoginFormDTO) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
}
