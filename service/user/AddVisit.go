package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func AddVisit(c *gin.Context) (interface{}, error) {
	var visit model.Visit
	err := c.BindJSON(&visit)
	if err != nil {
		return nil, err
	}

	userId := c.GetString("userId")
	if userId == "" {
		return nil, fmt.Errorf("user Id No specified in token")
	}

	visit.VisitId = uuid.New().String()
	visit.UserId = fmt.Sprint(userId)

	if err := validator.New().Struct(visit); err != nil {
		return nil, err
	}
	res, err := repository.CreateOne(&visit)
	return res, err

}
