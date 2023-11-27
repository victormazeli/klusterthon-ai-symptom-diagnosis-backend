package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/api/handlers"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/config"
	"gorm.io/gorm"
)

func BotRoute(cfg *config.Config, db *gorm.DB, r *gin.RouterGroup) {
	botHandler := handlers.BotHandler{
		Cfg: cfg,
		Db:  db,
	}

	r.POST("/chat", botHandler.HandleChat)
	//r.POST("/register", userHandler.Register)
}
