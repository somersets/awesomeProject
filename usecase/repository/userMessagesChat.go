package repository

import (
	"awesomeProject/domain"
	"awesomeProject/services/userChatService"
)

type UserChatMessagesRepository interface {
	SaveOne(message *domain.ChatMessageFormModel) error
	GetAllOfUser(userId int, recipientId int) (*[]userChatService.Message, error)
}
