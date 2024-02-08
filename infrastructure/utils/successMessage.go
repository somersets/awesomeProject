package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Success struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewSuccessResponse(c *gin.Context, statusCode int, message string) Success {
	logrus.Error(message)
	c.JSON(statusCode, Success{Message: message})
	return Success{
		message,
		statusCode,
	}
}
