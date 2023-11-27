package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/ApiResponse"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApiResponse.SendNotFound(c, "resource not found")
		c.Next()
	}
}
