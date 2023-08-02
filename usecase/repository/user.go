package repository

import (
	"awesomeProject/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateDisable(id int) (int, error)
	GetOneById(userId int) (*domain.UserDTO, error)
}
