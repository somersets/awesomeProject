package userChatRegistry

import (
	"awesomeProject/adapter/controller"
	"awesomeProject/adapter/repository"
	"awesomeProject/services/userChatService"
	"awesomeProject/usecase"
	"gorm.io/gorm"
)

func NewRegistry(db *gorm.DB) controller.UserChat {
	hub := userChatService.NewHub()
	u := usecase.NewUserChatUseCase(*hub, repository.NewDBRepository(db), repository.NewUserChatRepository(db))

	go u.RunUserChatHub()

	return controller.NewUserChatController(&u)
}
