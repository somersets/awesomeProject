package domain

import (
	"time"
)

type Token struct {
	ID           int        `gorm:"primary_key"`
	RefreshToken string     `gorm:"not null"`
	UserId       int        `gorm:"not null" gorm:"foreignKey:id"`
	CreatedAt    *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Token) TableName() string { return "tokens" }
