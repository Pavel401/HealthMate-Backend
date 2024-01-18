package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetBasicInfoByUserId(c *gin.Context) (interface{}, error) {
	userId := c.GetString("userId")
	if userId == "" {
		return nil, fmt.Errorf("invalid user Id")
	}
	item := model.Basic_Info{
		UserId: userId,
	}

	res, err := repository.GetOne(item)
	if err != nil {
		return nil, fmt.Errorf("Height format should be just 180  if 180cm integer")
	}

	//res.Bmi = res.Weight / res.Height
	return res, err
}
