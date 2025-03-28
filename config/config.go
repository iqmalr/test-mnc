package config

import "github.com/gin-gonic/gin"

func InitConfig() *gin.Engine {
	r := gin.Default()
	return r
}
