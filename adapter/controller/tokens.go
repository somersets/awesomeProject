package controller

import (
	"awesomeProject/infrastructure/utils"
	"awesomeProject/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tokensController struct {
	tokensUseCase usecase.TokensAuth
}

type TokensAuth interface {
	RefreshToken(c *gin.Context)
}

func NewTokensController(tuc *usecase.TokensAuth) TokensAuth {
	return &tokensController{*tuc}
}

func (tC *tokensController) RefreshToken(ctx *gin.Context) {
	token, err := ctx.Cookie("refreshToken")
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if len(token) == 0 {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, errors.New("empty cookie auth token").Error())
		return
	}

	refreshTokenResponse, err := tC.tokensUseCase.RefreshToken(token)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.SetCookie("refreshToken", refreshTokenResponse.RefreshToken, 30*24*60*60*1000, "", "localhost", true, true)
	ctx.JSON(http.StatusOK, refreshTokenResponse)
}
