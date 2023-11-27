package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/ApiResponse"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/config"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/database/models"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/dto"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Db  *gorm.DB
	Cfg *config.Config
}

func (u *AuthHandler) Login(c *gin.Context) {
	var loginInput dto.LoginDTO
	// deserialize JSON to struct
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		ApiResponse.SendBadRequest(c, err.Error())
		return
	}

	user := &models.User{}

	err := mgm.Coll(user).First(bson.M{"email": loginInput.Email}, user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ApiResponse.SendBadRequest(c, "invalid credentials")
			return
		} else {
			ApiResponse.SendInternalServerError(c, "an error occurred")
			return
		}

	}

	passwordMatch, err := utils.ComparePasswordAndHash(loginInput.Password, user.Password)
	if err != nil || !passwordMatch {
		ApiResponse.SendBadRequest(c, "invalid credentials")
		return
	}

	token := utils.GenerateToken(user.ID.Hex(), u.Cfg.Server.JwtKey)

	userToken := make(map[string]string)
	userToken["token"] = token
	//userToken["user_id"] = user.ID.Hex()

	ApiResponse.SendSuccess(c, "user logged in successfully", userToken)
	return
}

func (u *AuthHandler) Register(c *gin.Context) {
	var registerInput dto.RegisterDTO
	// deserialize JSON to struct
	if err := c.ShouldBindJSON(&registerInput); err != nil {
		ApiResponse.SendBadRequest(c, err.Error())
		return
	}

	user := &models.User{}

	e := mgm.Coll(user).First(bson.M{"email": registerInput.Email}, user)

	if e != nil {
		if e == mongo.ErrNoDocuments {
			hashPassword, e := utils.GenerateFromPassword(registerInput.Password)
			if e != nil {
				ApiResponse.SendInternalServerError(c, "an error occurred")
				return
			}

			newUser := &models.User{
				Email:    registerInput.Email,
				Password: hashPassword,
			}
			err := mgm.Coll(&models.User{}).Create(newUser)

			if err != nil {
				ApiResponse.SendInternalServerError(c, "an error occurred while creating user")
				return
			} else {
				ApiResponse.SendSuccess(c, "user created successfully", newUser)
				return
			}
		} else {
			ApiResponse.SendInternalServerError(c, "an error occurred")
			return
		}

	} else {
		ApiResponse.SendBadRequest(c, "user already exist")
		return
	}

}
