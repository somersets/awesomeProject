package controller

import (
	"awesomeProject/domain"
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userPhotoController struct {
	userUseCase usecase.UserPhoto
}

type UserPhoto interface {
	Create(ctx *gin.Context)
	SortOrder(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func (uc *userPhotoController) Delete(ctx *gin.Context) {
	userId, _ := utils.GetUserIdFromHeaderToken(ctx)
	var userPhotoDeleteForm domain.UserPhotoDeleteFormModel
	if err := ctx.ShouldBindJSON(&userPhotoDeleteForm); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	deletedUserPhoto, err := uc.userUseCase.Delete(userPhotoDeleteForm.ImageID, *userId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, deletedUserPhoto)
}

func (uc *userPhotoController) SortOrder(ctx *gin.Context) {
	userId, _ := utils.GetUserIdFromHeaderToken(ctx)
	var userPhotoOrderForm []domain.UserPhotoChangeOrderFormModel
	if err := ctx.ShouldBindJSON(&userPhotoOrderForm); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sortOrderErr := uc.userUseCase.SortOrder(&userPhotoOrderForm, *userId)
	if sortOrderErr != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, sortOrderErr.Error())
		return
	}

	ctx.JSON(http.StatusOK, utils.Success{
		Message: "OK",
		Status:  http.StatusOK,
	})
}

func (uc *userPhotoController) Update(ctx *gin.Context) {
	userId, _ := utils.GetUserIdFromHeaderToken(ctx)
	image, err := ctx.FormFile("image")
	imgId := ctx.PostForm("imageID")

	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if len(imgId) == 0 {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, errors.New("missing image_id").Error())
		return
	}

	imageId, err := strconv.Atoi(imgId)
	if err != nil || len(imgId) == 0 {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	newImage, createErr := uc.userUseCase.Update(*userId, imageId)
	if createErr != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, createErr.Error())
		return
	}

	uploadErr := ctx.SaveUploadedFile(image, fmt.Sprintf("infrastructure/assets/%s.png", newImage.PhotoName))
	if uploadErr != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, uploadErr.Error())
		return
	}
	ctx.JSON(http.StatusOK, newImage)
}

func (uc *userPhotoController) Create(ctx *gin.Context) {
	userId, _ := utils.GetUserIdFromHeaderToken(ctx)

	img, err := ctx.FormFile("image")
	orderImage := ctx.PostForm("orderImage")

	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	orderImg, err := strconv.Atoi(orderImage)
	if err != nil {
		return
	}

	createdImage, createErr := uc.userUseCase.Create(&domain.UserPhotoCreateFormModel{
		UserID: *userId,
		Order:  orderImg,
	}, img, ctx)

	if createErr != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, createErr.Error())
		return
	}

	ctx.JSON(http.StatusOK, createdImage)
}

func NewUserPhotoController(us *usecase.UserPhoto) UserPhoto {
	return &userPhotoController{*us}
}
