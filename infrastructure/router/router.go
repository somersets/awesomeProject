package router

import (
	"awesomeProject/adapter/controller"
	"awesomeProject/domain/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(c *controller.AppController) *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/refresh", c.Tokens.RefreshToken)
			auth.POST("/login", c.Auth.Login)
			auth.POST("/register", c.Auth.Register)
		}
		users := api.Group("users", middleware.VerifyAuth)
		{
			users.GET("/:id", c.User.GetUserById)
			users.POST("/create", c.User.CreateUser)
			users.PUT("/disable/:id", c.User.Disable)
		}
	}
	return router
}
