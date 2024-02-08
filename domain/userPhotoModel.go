package domain

import "time"

type UserPhoto struct {
	ID        int        `gorm:"primary_key" json:"id"`
	PhotoName string     `gorm:"not null" json:"photo"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserId    int        `json:"-"`
	Order     int        `json:"order"`
	Format    string     `json:"format"`
}

type UserPhotoCreateFormModel struct {
	UserID    int
	Order     int
	ImageName string
	Ext       string
}
type UserPhotoDeleteFormModel struct {
	ImageID int `json:"image_id" binding:"required"`
}
type UserPhotoChangeOrderFormModel struct {
	ImageID int `json:"image_id" binding:"required"`
	Order   int `json:"order" binding:"required"`
}
type UserPhotoUpdateFormModel struct {
	ImageID   int
	ImageName string
	UserID    int
}

func (UserPhoto) TableName() string { return "user-photos" }
