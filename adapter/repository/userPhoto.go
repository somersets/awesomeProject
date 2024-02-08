package repository

import (
	"awesomeProject/domain"
	"awesomeProject/usecase/repository"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type userPhotoRepository struct {
	db *gorm.DB
}

func (u userPhotoRepository) Delete(imageID int, userID int) (*domain.UserPhoto, error) {
	var deletedUserPhoto domain.UserPhoto
	if err := u.db.Model(&domain.UserPhoto{}).Where("id = ? AND user_id = ?", imageID, userID).First(&deletedUserPhoto).Error; err != nil {
		return nil, err
	}
	if err := u.db.Delete(&domain.UserPhoto{}, imageID).Error; err != nil {
		return nil, err
	}
	return &deletedUserPhoto, nil
}

func (u userPhotoRepository) Update(userPhotoForm *domain.UserPhotoUpdateFormModel) (*domain.UserPhoto, string, error) {
	var userPhoto *domain.UserPhoto
	fmt.Println(userPhotoForm.UserID)
	fmt.Println(userPhotoForm.ImageID)
	if err := u.db.Model(&domain.UserPhoto{}).Where("user_id = ? AND id = ?", userPhotoForm.UserID, userPhotoForm.ImageID).First(&userPhoto).Error; err != nil {
		return nil, "", err
	}

	oldImageName := userPhoto.PhotoName
	fmt.Println(oldImageName)

	userPhoto.PhotoName = userPhotoForm.ImageName
	u.db.Save(&userPhoto)

	return userPhoto, oldImageName, nil
}

func (u userPhotoRepository) Create(userFormModel *domain.UserPhotoCreateFormModel) (*domain.UserPhoto, error) {
	var userPhoto []domain.UserPhoto
	result := u.db.Model(&domain.UserPhoto{}).Where("user_id = ?", userFormModel.UserID).Find(&userPhoto)

	if result.RowsAffected >= 6 {
		return nil, errors.New("available max 6 photos")
	}

	newUserPhoto := domain.UserPhoto{
		PhotoName: userFormModel.ImageName,
		UserId:    userFormModel.UserID,
		Order:     userFormModel.Order,
		Format:    userFormModel.Ext,
	}
	if err := u.db.Model(&domain.UserPhoto{}).Create(&newUserPhoto).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return &newUserPhoto, nil
}

func (u userPhotoRepository) SortOrder(userPhotoChangeOrderForm *[]domain.UserPhotoChangeOrderFormModel, userID int) error {

	for _, image := range *userPhotoChangeOrderForm {
		var userPhoto *domain.UserPhoto
		if err := u.db.Model(&domain.UserPhoto{}).Where("user_id = ? AND id = ?", userID, image.ImageID).First(&userPhoto).Error; err != nil {
			return errors.New(err.Error())
		}

		userPhoto.Order = image.Order
		u.db.Save(&userPhoto)
	}

	return nil
}

func NewUserPhotoRepository(db *gorm.DB) repository.UserPhotoRepository {
	return &userPhotoRepository{db: db}
}
