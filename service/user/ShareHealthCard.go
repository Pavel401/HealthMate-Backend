package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func ShareHealthCard(c *gin.Context) (interface{}, error) {
	userId := c.GetString("userId")

	if userId == "" {
		return nil, fmt.Errorf("invalid user Id")
	}
	page := model.Public_Page{
		PageId:   uuid.New().String(),
		UserId:   userId,
		Code:     uuid.New().String()[1:5],
		ValidTil: time.Now().Add(time.Hour * 24),
	}

	res, err := repository.CreateOne(page)
	if err != nil {
		return nil, err
	}
	/*	res, err := repository.GetOne(item)
	 */
	return res, err
}
