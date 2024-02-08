package controller

import (
	"awesomeProject/domain"
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type authController struct {
	authUseCase usecase.Auth
}

type Auth interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
	Activate(c *gin.Context)
}

func NewAuthController(auc *usecase.Auth) Auth {
	return &authController{*auc}
}

func (ac *authController) Activate(ctx *gin.Context) {
	linkUUID := ctx.Param("uuid")
	if len(linkUUID) == 0 {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, errors.New("empty activation link uuid parameter").Error())
		return
	}

	loginResponse, err := ac.authUseCase.Activate(linkUUID)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.SetCookie("refreshToken", loginResponse.RefreshToken, 30*24*60*60*1000, "", "localhost", true, true)
	ctx.JSON(http.StatusOK, loginResponse)
}

func (ac *authController) Register(ctx *gin.Context) {
	var user *domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

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
	} else {
		ctx.SetCookie("refreshToken", loginResponse.RefreshToken, 30*24*60*60*1000, "", "localhost", true, true)
		ctx.JSON(http.StatusOK, loginResponse)
	}
}

func (ac *authController) Logout(ctx *gin.Context) {
	token, err := ctx.Cookie("refreshToken")
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if len(token) == 0 {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, errors.New("empty cookie auth token").Error())
		return
	}

	logoutError := ac.authUseCase.Logout(token)
	if logoutError != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, logoutError.Error())
		return
	}

	clearCookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		Secure:   true,
		MaxAge:   -1,
		Domain:   "localhost",
		HttpOnly: true,
	}

	ctx.SetCookie(clearCookie.Name, clearCookie.Value, clearCookie.MaxAge, clearCookie.Path, clearCookie.Domain, clearCookie.Secure, clearCookie.HttpOnly)

	utils.NewSuccessResponse(ctx, http.StatusBadRequest, "OK")
}
