package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllVisits(c *gin.Context) (interface{}, error) {
	userId := c.GetString("userId")

	if userId == "" {
		return nil, fmt.Errorf("invalid user Id")
	}
	visit := model.Visit{
		UserId: userId,
	}
	res, err := repository.GetAll(&visit, 1, 100, "", "")
	/*	res, err := repository.GetOne(item)
	 */
	return res, err
}
