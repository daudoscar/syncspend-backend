package routes

import (
	"net/http"
	"syncspend/controllers"
	"syncspend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	protected := router.Group("/protected")
	protected.Use(middleware.AuthenticateJWT())
	{
		// Protected Routes
	}

	return router
}
