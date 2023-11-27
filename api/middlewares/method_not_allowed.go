package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/ApiResponse"
)

func MethodNotAllowed() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApiResponse.SendMethodNotAllowedError(c, "method not allowed")
		c.Next()
	}
}
