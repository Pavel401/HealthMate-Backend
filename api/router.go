package api

import (
	handler "Hackathon_Backend/api/handlers"
	"Hackathon_Backend/api/middleware"
	_ "Hackathon_Backend/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.Default()) //allows all origin
	v1 := router.Group("/api/v1/")
	{

		users_api := v1.Group("/user/")
		{
			users_api.POST("/SignUp", handler.SignUp)
			users_api.POST("/login", handler.SignIn)
			users_api.POST("/basicInfo", middleware.TokenAuthMiddleware(), handler.AddBasicInfo)
			users_api.GET("/basicInfo", middleware.TokenAuthMiddleware(), handler.GetBasicInfo)
			users_api.PUT("/basicInfo", middleware.TokenAuthMiddleware(), handler.UpdateBasicInfo)
			users_api.POST("/measurement", middleware.TokenAuthMiddleware(), handler.AddMeasurement)
			users_api.GET("/measurement/:type", middleware.TokenAuthMiddleware(), handler.GetMeasurementByType)
			users_api.GET("/allMeasurement/:date", middleware.TokenAuthMiddleware(), handler.GetAllMeasurementByDate)
			users_api.GET("/allMeasurementOfDate/:date", middleware.TokenAuthMiddleware(), handler.GetAllMeasurementOfDate)
			users_api.POST("/upload", middleware.TokenAuthMiddleware(), handler.UploadFile)
			users_api.POST("/download", middleware.TokenAuthMiddleware(), handler.DownloadFile)
			users_api.POST("/visit", middleware.TokenAuthMiddleware(), handler.AddVisit)
			users_api.GET("/allVisit", middleware.TokenAuthMiddleware(), handler.GetAllVisit)
			users_api.GET("/share", middleware.TokenAuthMiddleware(), handler.ShareHealthCard)
			users_api.GET("/verify/:code", handler.ShareVerify)
			users_api.GET("/allHealthInsurance", middleware.TokenAuthMiddleware(), handler.GetAllHealthInsurance)
			users_api.POST("/addDonation", middleware.TokenAuthMiddleware(), handler.AddDonation)
			users_api.GET("/getAllDonations", handler.GetAllDonations)
			users_api.GET("/getAllDonationsUser", middleware.TokenAuthMiddleware(), handler.GetAllDonationsByUser)
			users_api.GET("/getAllPendingDonationsUser", handler.GetAllPendingDonations)
			users_api.GET("/clear", handler.ClearDownloads)
			users_api.PUT("/updateDonation/:id", handler.UpdateDonation)
		}

	}

	return router
}
