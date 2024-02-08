package middleware

import (
	"awesomeProject/adapter/repository"
	"awesomeProject/infrastructure/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func VerifyAuth(c *gin.Context) {
	header := c.GetHeader(authHeader)

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid auth header")
		return
	}
	_, err := repository.ValidateAccessToken(headerParts[1])
	if err != nil {
		utils.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
}
