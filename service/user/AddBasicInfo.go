package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func AddBasicInfo(c *gin.Context) (interface{}, error) {
	var info model.Basic_Info
	err := c.BindJSON(&info)
	if err != nil {
		return nil, err
	}

	userId, exists := c.Get("userId")
	if !exists {
		return nil, fmt.Errorf("user Id No specified in token")
	}

	info.BasicId = uuid.New().String()
	info.UserId = fmt.Sprint(userId)

	if err := validator.New().Struct(info); err != nil {
		return nil, err
	}
	res, err := repository.CreateOne(&info)
	return res, err

}
