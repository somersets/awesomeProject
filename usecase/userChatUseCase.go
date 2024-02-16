package usecase

import (
	"awesomeProject/domain"
	"awesomeProject/services/userChatService"
	"awesomeProject/usecase/repository"
)

type UserChat interface {
	RunUserChatHub()
	UserChatHub() *userChatService.Hub
	GetUserMessagesHistory(userId int, recipientId int) (*[]userChatService.Message, error)
	CreateChat()
}

type useUserChatUseCase struct {
	userChatMessagesRepository repository.UserChatMessagesRepository
	hubRepository              userChatService.Hub
	dbRepository               repository.DBRepository
}

func (u useUserChatUseCase) CreateChat() {
	//TODO implement me
	panic("implement me")
}

func NewUserChatUseCase(hubR userChatService.Hub, dbR repository.DBRepository, usrR repository.UserChatMessagesRepository) UserChat {
	return &useUserChatUseCase{hubRepository: hubR, dbRepository: dbR, userChatMessagesRepository: usrR}
}

func (u useUserChatUseCase) GetUserMessagesHistory(userId int, recipientId int) (*[]userChatService.Message, error) {
	messages, err := u.userChatMessagesRepository.GetAllOfUser(userId, recipientId)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (u useUserChatUseCase) UserChatHub() *userChatService.Hub {
	return &u.hubRepository
}
func (u useUserChatUseCase) RunUserChatHub() {
	for {
		select {
		case client := <-u.hubRepository.Register:
			u.hubRepository.RegisterNewClient(client)
		case client := <-u.hubRepository.Unregister:
			u.hubRepository.RemoveClient(client)
		case message := <-u.hubRepository.Broadcast:
			err := u.userChatMessagesRepository.SaveOne(&domain.ChatMessageFormModel{
				SenderId:    message.Sender.ID,
				RecipientId: message.Recipient.ID,
				Message:     message.Content,
			})
			if err != nil {
				return
			}
			u.hubRepository.HandleMessage(message)
		}
	}
}
