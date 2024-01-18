package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"Hackathon_Backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) (string, string, string, error) {
	userId := c.GetString("userId")

	if userId == "" {
		return "", "", "", fmt.Errorf("user Id No specified in token")
	}

	var file model.Aws_File

	err := c.BindJSON(&file)

	if err != nil {
		return "", "", "", fmt.Errorf("body is missing")
	}

	file.UserId = userId

	res, err := repository.GetOne(file)

	if err != nil {
		return "", "", "", err
	}

	err = utils.S3FileDownloader(res.AwsLink)
	if err != nil {
		fmt.Printf("%v", err)
		return "", "", "", err
	}

	err = utils.Decrypt("./recordsTemp/download/"+res.AwsLink, res.AwsLink+"."+res.Extension)
	if err != nil {
		fmt.Printf("%v", err)
		return "", "", "", err
	}
	/*c.FileAttachment("./recordsTemp/download/"+res.AwsLink+res.Extension, res.AwsLink+res.Extension)

	err = utils.DeleteFiles("../recordsTemp/download/"+res.AwsLink+res.Extension)
	if err != nil {
		return "","","", err
	}
	err = utils.DeleteFiles("./recordsTemp/download/"+res.AwsLink)
	if err != nil {
		return "","","", err
	}*/

	return "./recordsTemp/download/", res.AwsLink + res.Extension, "./recordsTemp/download/" + res.AwsLink, nil
}
