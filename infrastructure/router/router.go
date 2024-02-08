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
	router.MaxMultipartMemory = 8 << 20
	router.Use(func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)
		c.Next()
	})
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization", "credentials", "Site-Locale"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	api.Static("/users/photos/static", "infrastructure/assets")
	{
		auth := api.Group("/users/auth")
		{
			auth.GET("/refresh", c.Tokens.RefreshToken)
			auth.POST("/login", c.Auth.Login)
			auth.POST("/register", c.Auth.Register)
			auth.GET("/activate/:uuid", c.Auth.Activate)
			auth.GET("/logout", c.Auth.Logout, middleware.VerifyAuth)
		}
		users := api.Group("/users", middleware.VerifyAuth)
		{
			users.GET("/current", c.User.CurrentUser)
			users.GET("/:id", c.User.GetUserById)
			users.POST("/create", c.User.CreateUser)
			users.PUT("/current", c.User.UpdateUser)
			users.PUT("/disable/:id", c.User.Disable)
			users.GET("/orientations", c.User.GetSexOrientations)
		}
		photos := api.Group("/users/photos", middleware.VerifyAuth)
		{
			photos.POST("", c.UserPhoto.Create)
			photos.PATCH("", c.UserPhoto.Update)
			photos.DELETE("", c.UserPhoto.Delete)
			photos.PATCH("/order", c.UserPhoto.SortOrder)
		}
	}
	return router
}
