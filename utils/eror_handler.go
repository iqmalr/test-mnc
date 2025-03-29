package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

func HandleError(c *gin.Context, statusCode int, message string, err error, details ...interface{}) {
	response := ErrorResponse{
		Status:  statusCode,
		Message: message,
	}
	if err != nil {
		response.Error = err.Error()
	}
	if len(details) > 0 {
		response.Details = details[0]
	}
	c.JSON(statusCode, response)
}

func HandleValidationError(c *gin.Context, err error) {
	HandleError(c, 400, "Data tidak valid", err)
}

func HandleDatabaseError(c *gin.Context, operation string, err error) {
	message := fmt.Sprintf("Gagal %s data", operation)
	HandleError(c, 500, message, err)
}

func HandleNotFoundError(c *gin.Context, resource string) {
	message := fmt.Sprintf("%s tidak ditemukan", resource)
	HandleError(c, 404, message, nil)
}
