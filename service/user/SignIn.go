package user

import (
	"Hackathon_Backend/api/middleware"
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"Hackathon_Backend/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SignInUser(c *gin.Context) (interface{}, error) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		return nil, fmt.Errorf("input miss match")
	}
	if user.Email == "" || user.Password == "" {
		return nil, fmt.Errorf("email or password can't be empty")
	}

	userPassword := user.Password
	user.Password = ""
	realUser, err := repository.GetOne(user)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, fmt.Errorf("user not found")
		}

		return nil, err
	}

	err = utils.CheckPassword(userPassword, realUser.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "email or password is incorrect")
		return nil, fmt.Errorf("email or password is incorrect")
	}

	scope := []string{realUser.UserType}
	token, err := middleware.GenerateToken(realUser.UserId, realUser.UserType, scope)
	return token, err
}
