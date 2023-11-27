package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/ApiResponse"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/pkg/utils"
	"strings"
)

func Auth(jwtkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if len(strings.Split(bearerToken, " ")) == 2 {
			token := strings.Split(bearerToken, " ")[1]
			sub, err := utils.ValidateToken(token, jwtkey)
			if err != nil {
				ApiResponse.SendUnauthorized(c, err.Error())
				return
			} else {
				c.Set("user", sub)
				c.Next()
			}
		} else {
			ApiResponse.SendUnauthorized(c, "invalid token")
			return
		}
	}
}
