package domain

import "time"

type ChatMember struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserId    int
	ChatId    int
	User      User `foreignKey:"UserId"`
	Chat      Chat `foreignKey:"ChatId"`
}

func (ChatMember) TableName() string { return "user-chat-members" }
