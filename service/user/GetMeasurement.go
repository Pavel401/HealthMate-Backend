package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetMeasurement(c *gin.Context) (interface{}, error) {
	userId := c.GetString("userId")
	measurementType := c.Param("type")

	if userId == "" {
		return nil, fmt.Errorf("invalid user Id")
	}
	measurement := model.Measurement{
		UserId: userId,
		Type:   measurementType,
	}
	res, err := repository.GetAll(&measurement, 1, 20, "", "user_id = '"+userId+"' and measurements.type = '"+measurementType+"'")
	/*	res, err := repository.GetOne(item)
	 */
	return res, err
}
