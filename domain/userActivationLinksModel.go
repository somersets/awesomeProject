package domain

import "time"

type UserActivationLink struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Link      string     `json:"link"`
	ExpireAt  time.Time  `json:"expire_at"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserId    int
}

func (UserActivationLink) TableName() string { return "user-activation-links" }
