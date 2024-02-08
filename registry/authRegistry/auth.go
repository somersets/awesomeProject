package authRegistry

import (
	"awesomeProject/adapter/controller"
	"awesomeProject/adapter/repository"
	"awesomeProject/adapter/repository/auth"
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase"
	"gorm.io/gorm"
	"os"
)

func NewRegistry(db *gorm.DB) controller.Auth {
	u := usecase.NewAuthUseCase(
		auth.NewAuthRepository(db),
		repository.NewTokensRepository(db),
		repository.NewDBRepository(db),
		utils.NewMailService(
			os.Getenv("facelessmay_mail_IDENTITY"),
			os.Getenv("facelessmay_mail_EMAILFROM"),
			os.Getenv("facelessmay_mail_PASSWORDFROM"),
			os.Getenv("EMAILTO_1"),
		),
	)

	return controller.NewAuthController(&u)
}
