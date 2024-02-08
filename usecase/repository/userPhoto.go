package repository

import "awesomeProject/domain"

type UserPhotoRepository interface {
	Create(userPhotoForm *domain.UserPhotoCreateFormModel) (*domain.UserPhoto, error)
	SortOrder(userPhotoChangeOrderForm *[]domain.UserPhotoChangeOrderFormModel, userID int) error
	Update(userPhotoForm *domain.UserPhotoUpdateFormModel) (*domain.UserPhoto, string, error)
	Delete(imageID int, userID int) (*domain.UserPhoto, error)
}
