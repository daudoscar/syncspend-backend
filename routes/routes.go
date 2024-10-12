package routes

import (
	"syncspend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// User routes
	router.POST("/register", controllers.RegisterUser)

	return router
}
