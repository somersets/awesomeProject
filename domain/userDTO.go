package domain

import (
	"time"
)

type UserDTO struct {
	ID             int                        `json:"id"`
	FirstName      string                     `json:"first_name"`
	LastName       string                     `json:"last_name"`
	Patronymic     string                     `json:"patronymic"`
	Email          string                     `json:"email"`
	Phone          int                        `json:"phone"`
	Birthday       *time.Time                 `json:"birthday"`
	Activated      bool                       `json:"activated"`
	Gender         string                     `json:"gender"`
	CreatedAt      *time.Time                 `json:"created_at"`
	UpdatedAt      *time.Time                 `json:"updated_at"`
	DisabledAt     *time.Time                 `json:"disabled_at"`
	DisabledUntil  *time.Time                 `json:"disabled_until"`
	SexOrientation *UserSexualOrientationType `json:"sex_orientation"`
	Photos         *[]UserPhoto               `json:"photos"`
}

func NewUserDTO(user *User) *UserDTO {
	return &UserDTO{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Patronymic:     user.Patronymic,
		Email:          user.Email,
		Phone:          user.Phone,
		Birthday:       user.Birthday,
		Activated:      user.Activated,
		Gender:         user.Gender,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DisabledAt:     user.DisabledAt,
		DisabledUntil:  user.DisabledUntil,
		SexOrientation: user.SexOrientation,
		Photos:         user.Photos,
	}
}
