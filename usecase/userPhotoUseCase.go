package usecase

import (
	"awesomeProject/domain"
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type UserPhoto interface {
	SortOrder(userPhotoChangeOrderForm *[]domain.UserPhotoChangeOrderFormModel, userID int) error
	Create(userPhotoCreateForm *domain.UserPhotoCreateFormModel, image *multipart.FileHeader, ctx *gin.Context) (*domain.UserPhoto, error)
	Update(userID int, imageID int) (*domain.UserPhoto, error)
	Delete(imageID int, userID int) (*domain.UserPhoto, error)
}

type userPhotoUseCase struct {
	userPhotoRepository repository.UserPhotoRepository
	dbRepository        repository.DBRepository
}

func (u userPhotoUseCase) Delete(imageID int, userID int) (*domain.UserPhoto, error) {
	deletedUserPhoto, err := u.userPhotoRepository.Delete(imageID, userID)
	if err != nil {
		return nil, err
	}

	err = os.Remove(fmt.Sprintf("infrastructure/assets/%s%s", deletedUserPhoto.PhotoName, deletedUserPhoto.Format))
	if err != nil {
		return nil, err
	}

	return deletedUserPhoto, nil
}

func (u userPhotoUseCase) SortOrder(userPhotoChangeOrderForm *[]domain.UserPhotoChangeOrderFormModel, userID int) error {
	err := u.userPhotoRepository.SortOrder(userPhotoChangeOrderForm, userID)

	if err != nil {
		return err
	}

	return nil
}

func (u userPhotoUseCase) Create(userPhotoCreateForm *domain.UserPhotoCreateFormModel, image *multipart.FileHeader, ctx *gin.Context) (*domain.UserPhoto, error) {
	photoName := uuid.Must(uuid.NewRandom()).String()
	// TODO compress image
	imageExtension := filepath.Ext(image.Filename)

	imageCreateForm := &domain.UserPhotoCreateFormModel{
		UserID:    userPhotoCreateForm.UserID,
		Order:     userPhotoCreateForm.Order,
		ImageName: photoName,
		Ext:       imageExtension,
	}
	userPhotoCreateForm.ImageName = photoName
	userPhotoCreateForm.Ext = imageExtension

	newImage, err := u.userPhotoRepository.Create(imageCreateForm)
	if err != nil {
		return nil, err
	}

	uploadErr := ctx.SaveUploadedFile(image, fmt.Sprintf("infrastructure/assets/%s.png", newImage.PhotoName))
	if uploadErr != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, uploadErr.Error())
		return nil, uploadErr
	}

	return newImage, nil
}

func (u userPhotoUseCase) Update(userID int, imageID int) (*domain.UserPhoto, error) {
	photoName := uuid.Must(uuid.NewRandom()).String()

	userPhoto, oldImageName, err := u.userPhotoRepository.Update(&domain.UserPhotoUpdateFormModel{
		ImageID:   imageID,
		ImageName: photoName,
		UserID:    userID,
	})

	if err != nil {
		return nil, err
	}

	if len(oldImageName) > 0 {
		err := os.Remove(fmt.Sprintf("infrastructure/assets/%s%s", oldImageName, userPhoto.Format))
		if err != nil {
			return nil, err
		}
	}

	return userPhoto, nil
}

func NewUserPhotoUseCase(upR repository.UserPhotoRepository, dbR repository.DBRepository) UserPhoto {
	return &userPhotoUseCase{
		userPhotoRepository: upR,
		dbRepository:        dbR,
	}
}
