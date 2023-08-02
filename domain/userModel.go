package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            int            `gorm:"primary_key" json:"id"`
	FirstName     string         `gorm:"not null" json:"first_name" binding:"required"`
	LastName      string         `gorm:"not null" json:"last_name" binding:"required"`
	Patronymic    string         `json:"patronymic"`
	Password      *string        `gorm:"not null" json:"password,omitempty" binding:"required"`
	Email         string         `gorm:"not null" json:"email" binding:"required"`
	Phone         int            `gorm:"not bull" json:"phone" binding:"required"`
	Birthday      *time.Time     `json:"birthday"`
	Gender        string         `gorm:"not null" gorm:"default:false" json:"gender" binding:"required"`
	CreatedAt     *time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DisabledAt    *time.Time     `json:"disabled_at"`
	DisabledUntil *time.Time     `json:"disabled_until"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Token         []Token        `gorm:"foreignKey:UserId" json:"omitempty"`
}

func (User) TableName() string { return "users" }
