package routes

import (
	"test-mnc/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/merchants", controllers.GetAllMerchants)
	r.POST("/installment", controllers.CreateInstallment)
	r.GET("/installment", controllers.GetAllInstallment)
	r.GET("/payments", controllers.GetAllPayments)
	r.POST("/payments", controllers.CreatePayment)
	r.GET("/recap", controllers.GetInstallmentRecap)
	return r
}
