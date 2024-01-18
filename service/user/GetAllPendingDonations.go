package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"github.com/gin-gonic/gin"
)

func GetAllPendingDonationsByUser(c *gin.Context) (interface{}, error) {
	donation := model.Drug_Donation{}
	res, err := repository.GetAll(&donation, 1, 100, "", "status = 'Requested'")
	if err != nil {
		return nil, err
	}

	/*	res, err := repository.GetOne(item)
	 */
	return res, err
}
