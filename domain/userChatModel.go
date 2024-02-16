package domain

import "time"

type Chat struct {
	ID                  int                   `gorm:"primary_key" json:"id"`
	ChatName            string                `json:"chat_name"`
	ChatMessagesHistory *[]ChatMessageHistory `gorm:"foreignKey:ChatId"`
}

type SenderAndRecipientInfoMessage struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	LastName string     `json:"last_name"`
	Date     *time.Time `json:"date"`
}

type ChatMessageFormModel struct {
	SenderId    int    `json:"sender_id" binding:"required"`
	RecipientId int    `json:"recipient_id" binding:"required"`
	Message     string `json:"message" binding:"required"`
}

func (Chat) TableName() string { return "user-chats" }
