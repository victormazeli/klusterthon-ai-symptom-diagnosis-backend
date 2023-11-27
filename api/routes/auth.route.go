package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/api/handlers"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/config"
	"gorm.io/gorm"
)

func AuthRoute(cfg *config.Config, db *gorm.DB, r *gin.RouterGroup) {
	authHandler := handlers.AuthHandler{
		Cfg: cfg,
		Db:  db,
	}

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
}
