package routes

import (
	"test-mnc/config"
	"test-mnc/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// r.GET("/users", controllers.GetAllUsers)
	r.GET("/merchants", controllers.GetAllMerchants)
	r.POST("/installment", controllers.CreateInstallment)
	r.GET("/installment", controllers.GetAllInstallment)
	r.GET("/payments", controllers.GetAllPayments)
	r.POST("/payments", controllers.CreatePayment)
	r.GET("/recap", controllers.GetInstallmentRecap)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	protected := r.Group("/")
	protected.Use(config.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetAllUsers)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
