package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllMeasurement(c *gin.Context) (interface{}, error) {
	userId := c.GetString("userId")
	date := c.Param("date")
	if date == "" {
		return nil, fmt.Errorf("invalid Date")
	}

	if userId == "" {
		return nil, fmt.Errorf("invalid user Id")
	}
	measurement := model.Measurement{
		UserId: userId,
	}
	res, err := repository.GetAll(&measurement, 1, 100, "", "created_at >= '"+date+"'")
	/*	res, err := repository.GetOne(item)
	 */
	return res, err
}
