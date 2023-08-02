package authRegistry

import (
	"awesomeProject/adapter/controller"
	"awesomeProject/adapter/repository"
	"awesomeProject/adapter/repository/auth"
	"awesomeProject/usecase"
	"gorm.io/gorm"
)

func NewRegistry(db *gorm.DB) controller.Auth {
	u := usecase.NewAuthUseCase(auth.NewAuthRepository(db), repository.NewTokensRepository(db), repository.NewDBRepository(db))

	return controller.NewAuthController(&u)
}
