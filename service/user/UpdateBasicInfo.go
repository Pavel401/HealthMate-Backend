package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UpdateBasicInfo(c *gin.Context) (interface{}, error) {
	var info model.Basic_Info
	userId := c.GetString("userId")

	if userId == "" {
		return nil, fmt.Errorf("invalid menu id")
	}
	info.UserId = userId

	itemExist, err := repository.GetOne(&info)

	if err != nil && itemExist.UserId == userId {
		return nil, err
	}

	err = c.BindJSON(&info)
	if err != nil {
		return nil, err
	}
	info.BasicId = itemExist.BasicId
	res, err := repository.UpdateOne(&info)
	return res, err
}
