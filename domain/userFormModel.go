package domain

import "time"

type UserProfileUpdateForm struct {
	FirstName        string     `json:"first_name" binding:"required"`
	LastName         string     `json:"last_name" binding:"required"`
	Email            string     `json:"email" binding:"required"`
	Phone            int        `json:"phone" binding:"required"`
	Birthday         *time.Time `json:"birthday" binding:"required"`
	Gender           string     `json:"gender" binding:"required"`
	SexOrientationID int        `json:"sex_orientation_id" binding:"required"`
}
