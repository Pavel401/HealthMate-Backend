package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllDonationsByUser(c *gin.Context) (interface{}, error) {
	userId := c.GetString("userId")

	if userId == "" {
		return nil, fmt.Errorf("invalid user Id")
	}

	donation := model.Drug_Donation{
		UserId: userId,
	}
	res, err := repository.GetAll(&donation, 1, 100, "", "")
	if err != nil {
		return nil, err
	}
	/*	res, err := repository.GetOne(item)
	 */
	return res, err
}

//user_id = '"+userId+"'
