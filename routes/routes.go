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
			plans.GET("/:id", controllers.GetPlanById)
			plans.POST("", controllers.CreatePlan)
			plans.PUT("/:id", controllers.UpdatePlan)
			plans.DELETE("/:id", controllers.DeletePlan)
			plans.PUT("/:id/recover", controllers.RecoverPlan)

			plans.POST("/join", controllers.JoinPlan)
			plans.DELETE("/leave", controllers.LeavePlan)
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
