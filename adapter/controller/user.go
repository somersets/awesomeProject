package controller

import (
	"awesomeProject/domain"
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userController struct {
	userUseCase usecase.User
}

type User interface {
	Disable(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
}

func NewUserController(us *usecase.User) User {
	return &userController{*us}
}

func (uc *userController) GetUserById(ctx *gin.Context) {
	userIdParam := ctx.Param("id")
	if len(userIdParam) == 0 {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, errors.New("empty userId parameter").Error())
		return
	}

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := uc.userUseCase.GetOneById(userId)
	ctx.JSON(http.StatusOK, user)
}

func (uc *userController) Disable(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if len(idParam) == 0 {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, errors.New("id is not found").Error())
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	u, err := uc.userUseCase.Disable(id)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, u)
}

func (uc *userController) CreateUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hashPass, _ := utils.HashPassword(*user.Password)
	user.Password = &hashPass

	u, err := uc.userUseCase.Create(&user)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userDto := domain.NewUserDTO(u)

	ctx.JSON(http.StatusCreated, userDto)
}
