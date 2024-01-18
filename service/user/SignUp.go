package user

import (
	"Hackathon_Backend/api/middleware"
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"Hackathon_Backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func SignUp(c *gin.Context) (interface{}, error) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		return nil, err
	}
	user.UserId = uuid.New().String()
	if user.Password == "" {
		return nil, fmt.Errorf("password can't be empty")
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	if err := validator.New().Struct(user); err != nil {
		return nil, err
	}
	_, err = repository.CreateOne(&user)
	if err != nil {
		return nil, err
	}

	scope := []string{"normal"}
	token, err := middleware.GenerateToken(user.UserId, "normal", scope)

	return token, err
}
