package utils

import (
	"awesomeProject/adapter/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetUserIdFromHeaderToken(ctx *gin.Context) (*int, error) {
	bearerAuth := ctx.GetHeader("Authorization")
	if len(bearerAuth) == 0 {
		authErr := NewErrorResponse(ctx, http.StatusBadRequest, errors.New("empty authorization header").Error())
		return nil, errors.New(authErr.Message)
	}

	splitToken := strings.Split(bearerAuth, "Bearer ")
	token := splitToken[1]
	userId, err := repository.ValidateAccessToken(token)
	if err != nil {
		authErr := NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return nil, errors.New(authErr.Message)
	}

	return userId, nil
}
