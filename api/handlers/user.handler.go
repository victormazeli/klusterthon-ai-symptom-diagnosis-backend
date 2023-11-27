package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/ApiResponse"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/config"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/database/models"
	"gorm.io/gorm"
)

type UserHandler struct {
	Db  *gorm.DB
	Cfg *config.Config
}

func (u *UserHandler) GetCurrentLoggedInUser(c *gin.Context) {
	userId, _ := c.Get("user")

	user := &models.User{}

	err := mgm.Coll(user).FindByID(userId, user)

	if err != nil {
		ApiResponse.SendInternalServerError(c, "internal server error")
	}

	ApiResponse.SendSuccess(c, "user fetched successfully", user)
	return

}
