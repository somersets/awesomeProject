package router

import (
	"awesomeProject/adapter/controller"
	"awesomeProject/domain/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(c *controller.AppController) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization", "credentials"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
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
			users.GET("/logout", c.Auth.Logout)
			users.GET("/:id", c.User.GetUserById)
			users.POST("/create", c.User.CreateUser)
			users.PUT("/disable/:id", c.User.Disable)
		}
	}
	return router
}
