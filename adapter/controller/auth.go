package controller

import (
	"awesomeProject/domain"
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authController struct {
	authUseCase usecase.Auth
}

type Auth interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

func NewAuthController(auc *usecase.Auth) Auth {
	return &authController{*auc}
}

func (ac *authController) Register(ctx *gin.Context) {
	var user *domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hashPass, _ := utils.HashPassword(*user.Password)
	user.Password = &hashPass

	registerResponseDTO, err := ac.authUseCase.Register(user)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.SetCookie("refreshToken", registerResponseDTO.RefreshToken, 30*24*60*60*1000, "", "localhost", true, true)
	ctx.JSON(http.StatusCreated, registerResponseDTO)
}

func (ac *authController) Login(ctx *gin.Context) {
	var loginForm *domain.LoginFormDTO
	if err := ctx.ShouldBindJSON(&loginForm); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if loginResponse, err := ac.authUseCase.Login(loginForm); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	} else {
		ctx.SetCookie("refreshToken", loginResponse.RefreshToken, 30*24*60*60*1000, "", "*", true, true)
		ctx.JSON(http.StatusOK, loginResponse)
	}
}
