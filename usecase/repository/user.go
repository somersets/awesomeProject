package repository

import (
	"awesomeProject/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) (*domain.UserDTO, error)
	UpdateDisable(id int) (int, error)
	GetOneById(userId int) (*domain.UserDTO, error)
	UpdateUser(userFormUpdate *domain.UserProfileUpdateForm, userId int) (*domain.UserDTO, error)
	GetSexOrientations() ([]*domain.UserSexualOrientationType, error)
}
