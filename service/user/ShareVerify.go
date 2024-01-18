package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ShareVerify(c *gin.Context) (interface{}, error) {
	code := c.Param("code")
	if code == "" {
		return nil, fmt.Errorf("invalid Code")
	}

	healthCard := model.Public_Page{
		Code: code,
	}

	record, err := repository.GetOne(healthCard)
	if err != nil {
		return nil, err
	}

	user := model.User{
		UserId: record.UserId,
	}

	basic := model.Basic_Info{
		UserId: record.UserId,
	}

	basicInfo, err := repository.GetOne(basic)
	if err != nil {
		return nil, fmt.Errorf("Height format should be just 180  if 180cm integer")
	}

	visit := model.Visit{
		UserId: user.UserId,
	}
	res, err := repository.GetAll(&visit, 1, 100, "", "")

	profile := model.SharedProfile{
		BasicInfo: basicInfo,
		Visits:    res.Results,
	}
	/*	res, err := repository.GetOne(item)
	 */
	return profile, err
}
