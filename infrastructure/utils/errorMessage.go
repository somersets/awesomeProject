package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) Error {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
	return Error{
		message,
		statusCode,
	}
}
