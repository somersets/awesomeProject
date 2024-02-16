package domain

import "time"

type ChatMessageHistory struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Text      string     `gorm:"not null" json:"text"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ChatId    int        `gorm:"not null" json:"chat_id"`
	SenderId  int        `gorm:"not null" json:"sender_id"`
}

func (ChatMessageHistory) TableName() string {
	return "user-chat-messages-history"
}
