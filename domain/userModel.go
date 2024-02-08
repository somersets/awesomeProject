package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID               int                `gorm:"primary_key" json:"id"`
	FirstName        string             `json:"first_name"`
	LastName         string             `json:"last_name"`
	Patronymic       string             `json:"patronymic"`
	Password         *string            `gorm:"not null" json:"password,omitempty" binding:"required"`
	Email            string             `gorm:"not null" json:"email" binding:"required"`
	Phone            int                `gorm:"not null" json:"phone" binding:"required"`
	Birthday         *time.Time         `json:"birthday"`
	Gender           string             `json:"gender"`
	CreatedAt        *time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        *time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	DisabledAt       *time.Time         `json:"disabled_at"`
	DisabledUntil    *time.Time         `json:"disabled_until"`
	Activated        bool               `json:"activated" gorm:"default:false"`
	DeletedAt        gorm.DeletedAt     `json:"deleted_at"`
	Token            []Token            `gorm:"foreignKey:UserId" json:"omitempty"`
	Photos           *[]UserPhoto       `gorm:"foreignKey:UserId"`
	ActivationLink   UserActivationLink `gorm:"foreignKey:UserId"`
	SexOrientationID *int               `gorm:"foreignKey:id" json:"sex_orientation_id"`
	SexOrientation   *UserSexualOrientationType
}

func (User) TableName() string { return "users" }
