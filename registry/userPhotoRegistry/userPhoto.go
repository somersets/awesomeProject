package userPhotoRegistry

import (
	"awesomeProject/adapter/controller"
	"awesomeProject/adapter/repository"
	"awesomeProject/usecase"
	"gorm.io/gorm"
)

func NewRegistry(db *gorm.DB) controller.UserPhoto {
	u := usecase.NewUserPhotoUseCase(
		repository.NewUserPhotoRepository(db),
		repository.NewDBRepository(db),
	)

	return controller.NewUserPhotoController(&u)
}
