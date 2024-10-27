package routes

import (
	"syncspend/controllers"
	"syncspend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	protected := router.Group("/protected")
	protected.Use(middleware.AuthenticateJWT())
	{
		plans := protected.Group("/plans")
		{
			plans.POST("", controllers.CreatePlan)
			plans.PUT("/:id", controllers.UpdatePlan)
			plans.GET("/:id", controllers.CreatePlan)
			plans.DELETE("/:id", controllers.CreatePlan)
			plans.GET("", controllers.CreatePlan)
		}

		transaksi := protected.Group("/transaksi")
		{
			transaksi.POST("", controllers.CreateTransaksi)
			transaksi.PUT("/:id", controllers.CreateTransaksi)
			transaksi.GET("/:id", controllers.CreateTransaksi)
			transaksi.DELETE("/:id", controllers.CreateTransaksi)
			transaksi.GET("", controllers.CreateTransaksi)
		}

		users := protected.Group("/users")
		{
			users.GET("/profile", controllers.GetUserByID)
			users.PUT("/profile", controllers.UpdateUser)
		}
	}

	return router
}
