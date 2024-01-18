package handler

import (
	"Hackathon_Backend/service/user"
	"Hackathon_Backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	res, err := user.SignUp(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func SignIn(c *gin.Context) {
	res, err := user.SignInUser(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func AddBasicInfo(c *gin.Context) {
	res, err := user.AddBasicInfo(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetBasicInfo(c *gin.Context) {
	res, err := user.GetBasicInfoByUserId(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func UpdateBasicInfo(c *gin.Context) {
	res, err := user.UpdateBasicInfo(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func AddMeasurement(c *gin.Context) {
	res, err := user.AddMeasurement(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetMeasurementByType(c *gin.Context) {
	res, err := user.GetMeasurement(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllMeasurementByDate(c *gin.Context) {
	res, err := user.GetAllMeasurement(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllMeasurementOfDate(c *gin.Context) {
	res, err := user.GetAllMeasurementByDate(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func UploadFile(c *gin.Context) {
	res, err := user.UploadFile(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func DownloadFile(c *gin.Context) {
	file, name, _, err := user.DownloadFile(c)
	fmt.Printf("\n%v\n", err)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		c.AbortWithStatusJSON(404, "Something went wrong")
	}
	c.FileAttachment(file, name)
	c.JSON(200, "ok")
}

func AddVisit(c *gin.Context) {
	res, err := user.AddVisit(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllVisit(c *gin.Context) {
	res, err := user.GetAllVisits(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func ShareHealthCard(c *gin.Context) {
	res, err := user.ShareHealthCard(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllHealthInsurance(c *gin.Context) {
	res, err := user.GetAllHealthInsurance(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func AddDonation(c *gin.Context) {
	res, err := user.AddDonation(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllDonations(c *gin.Context) {
	res, err := user.GetAllDonations(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllDonationsByUser(c *gin.Context) {
	res, err := user.GetAllDonationsByUser(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}
func ShareVerify(c *gin.Context) {
	res, err := user.ShareVerify(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllPendingDonations(c *gin.Context) {
	res, err := user.GetAllPendingDonationsByUser(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func ClearDownloads(c *gin.Context) {
	res, err := user.Clear(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func UpdateDonation(c *gin.Context) {
	res, err := user.UpdateDonation(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}
