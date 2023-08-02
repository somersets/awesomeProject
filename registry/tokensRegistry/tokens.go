package tokensRegistry

import (
	"awesomeProject/adapter/controller"
	"awesomeProject/adapter/repository"
	"awesomeProject/usecase"
	"gorm.io/gorm"
)

func NewRegistry(db *gorm.DB) controller.TokensAuth {
	tuc := usecase.NewTokensUseCase(
		repository.NewTokensRepository(db),
		repository.NewDBRepository(db),
	)

	return controller.NewTokensController(&tuc)
}
