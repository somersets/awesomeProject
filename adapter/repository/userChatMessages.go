package repository

import (
	"awesomeProject/domain"
	"awesomeProject/services/userChatService"
	"awesomeProject/usecase/repository"
	"gorm.io/gorm"
	"strconv"
)

type userChatRepository struct {
	db *gorm.DB
}

func NewUserChatRepository(db *gorm.DB) repository.UserChatMessagesRepository {
	return &userChatRepository{db: db}
}

func (u userChatRepository) GetAllOfUser(userId int, recipientId int) (*[]userChatService.Message, error) {
	var messagesHistory []domain.ChatMessageHistory
	if err := u.db.Model(&domain.ChatMessageHistory{}).Where("sender_id = ? AND recipient_id = ? OR recipient_id = ? AND sender_id = ?", userId, recipientId, userId, recipientId).Find(&messagesHistory).Error; err != nil {
		return nil, err
	}

	var user *domain.User
	if err := u.db.Model(&domain.User{}).Find(&user, userId).Error; err != nil {
		return nil, err
	}
	var recipient *domain.User
	if err := u.db.Model(&domain.User{}).Find(&recipient, recipientId).Error; err != nil {
		return nil, err
	}
	var messages = make([]userChatService.Message, len(messagesHistory))

	for index, message := range messagesHistory {
		messages[index] = userChatService.Message{
			Type: "message",
			Sender: domain.SenderAndRecipientInfoMessage{
				ID:       message.SenderId,
				Name:     user.FirstName,
				LastName: user.LastName,
				Date:     message.CreatedAt,
			},
			Recipient: domain.SenderAndRecipientInfoMessage{
				ID:       0,
				Name:     recipient.FirstName,
				LastName: recipient.LastName,
				Date:     message.CreatedAt,
			},
			Content: message.Text,
			ID:      strconv.Itoa(message.ID),
		}
	}

	return &messages, nil
}

func (u userChatRepository) SaveOne(message *domain.ChatMessageFormModel) error {
	savedMessage := &domain.ChatMessageHistory{
		Text:     message.Message,
		SenderId: message.SenderId,
	}
	result := u.db.Model(&domain.ChatMessageHistory{}).Create(&savedMessage)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
