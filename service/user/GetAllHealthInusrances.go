package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"github.com/gin-gonic/gin"
)

func GetAllHealthInsurance(c *gin.Context) (interface{}, error) {

	hi := model.Health_insurance{}
	res, err := repository.GetAll(&hi, 1, 100, "", "")
	/*	res, err := repository.GetOne(item)
	 */
	return res, err
}
