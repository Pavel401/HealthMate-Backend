package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func AddDonation(c *gin.Context) (interface{}, error) {
	var donation model.Drug_Donation
	err := c.BindJSON(&donation)
	if err != nil {
		return nil, err
	}

	userId, exists := c.Get("userId")
	if !exists {
		return nil, fmt.Errorf("user Id No specified in token")
	}

	donation.DonationId = uuid.New().String()
	donation.UserId = fmt.Sprint(userId)
	donation.Status = "Requested"

	if err := validator.New().Struct(donation); err != nil {
		return nil, err
	}
	res, err := repository.CreateOne(&donation)
	return res, err

}
