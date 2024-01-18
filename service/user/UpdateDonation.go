package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UpdateDonation(c *gin.Context) (interface{}, error) {
	var info model.Drug_Donation
	dId := c.Param("id")

	if dId == "" {
		return nil, fmt.Errorf("invalid id")
	}
	info.DonationId = dId

	itemExist, err := repository.GetOne(&info)

	if err != nil && itemExist.DonationId == dId {
		return nil, err
	}

	err = c.BindJSON(&info)
	if err != nil {
		return nil, err
	}
	info.UserId = itemExist.UserId
	info.DonationId = itemExist.DonationId
	res, err := repository.UpdateOne(&info)
	return res, err
}
