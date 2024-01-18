package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func AddMeasurement(c *gin.Context) (interface{}, error) {
	var measurement model.Measurement
	err := c.BindJSON(&measurement)
	if err != nil {
		return nil, err
	}

	userId := c.GetString("userId")
	if userId == "" {
		return nil, fmt.Errorf("user Id No specified in token")
	}

	measurement.MeasurementId = uuid.New().String()
	measurement.UserId = fmt.Sprint(userId)

	if err := validator.New().Struct(measurement); err != nil {
		return nil, err
	}
	res, err := repository.CreateOne(&measurement)
	return res, err

}
