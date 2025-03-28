package routes

import (
	"test-mnc/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/merchants", controllers.GetAllMerchants)
	r.POST("/transactions", controllers.CreateTransaction)
	return r
}
