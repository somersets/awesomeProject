package domain

import (
	"time"
)

type UserDTO struct {
	ID            int        `json:"id"`
	FirstName     string     `json:"first_name" binding:"required"`
	LastName      string     `json:"last_name" binding:"required"`
	Patronymic    string     `json:"patronymic"`
	Email         string     `json:"email" binding:"required"`
	Phone         int        `json:"phone" binding:"required"`
	Birthday      *time.Time `json:"birthday"`
	Gender        string     `json:"gender" binding:"required"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DisabledAt    *time.Time `json:"disabled_at"`
	DisabledUntil *time.Time `json:"disabled_until"`
}

func NewUserDTO(user *User) *UserDTO {
	return &UserDTO{
		ID:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Patronymic:    user.Patronymic,
		Email:         user.Email,
		Phone:         user.Phone,
		Birthday:      user.Birthday,
		Gender:        user.Gender,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		DisabledAt:    user.DisabledAt,
		DisabledUntil: user.DisabledUntil,
	}
}
