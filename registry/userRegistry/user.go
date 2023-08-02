package userRegistry

import (
	"awesomeProject/adapter/controller"
	"awesomeProject/adapter/repository"
	"awesomeProject/usecase"
	"gorm.io/gorm"
)

func NewRegistry(db *gorm.DB) controller.User {
	u := usecase.NewUserUseCase(
		repository.NewUserRepository(db),
		repository.NewDBRepository(db),
	)

	return controller.NewUserController(&u)
}
