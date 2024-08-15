package api

import (
	"api/api/handler"
	"api/api/middleware"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Voting service
// @version 1.0
// @description Voting service
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		log.Fatal("casbin error load policy: ", err)
		panic(err)
	}

	auth := r.Group("/auth")
	auth.Use(middleware.NewAuth(ca))
	// Auth
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/validateToken", h.ValidateToken)
		auth.POST("/refreshToken", h.RefreshToken)
		auth.POST("/validateEmail", h.ValidateEmail)
		auth.POST("/verifyCode", h.VerificationCode)
	}

	user := r.Group("/users")
	user.Use(middleware.NewAuth(ca))
	// Users
	{
		user.GET("/profile/:user_id", h.GetProfileInfo)
		user.PUT("/updateProfile", h.UpdateProfile)
		user.DELETE("/deleteProfile/:user_id", h.DeleteProfile)
		user.POST("/change-password", h.ChangePassword)
	}

	analytics := r.Group("/analytics")
	// analytics.Use(middleware.NewAuth(ca))
	{
		// Medical Records
		analytics.POST("/medical-record", h.AddMedicalRecord)
		analytics.GET("/medical-records", h.GetMedicalRecords)
		analytics.GET("/medical-record/:id", h.GetMedicalRecordById)
		analytics.PUT("/medical-record", h.UpdateMedicalRecord)
		analytics.DELETE("/medical-record/:id", h.DeleteMedicalRecord)
		analytics.GET("/medical-records/list", h.ListMedicalRecords)

		// Lifestyle Data
		analytics.POST("/lifestyle-data", h.AddLifestyleData)
		analytics.GET("/lifestyle-data", h.GetLifestyleData)
		analytics.GET("/lifestyle-data/:id", h.GetLifestyleDataById)
		analytics.PUT("/lifestyle-data", h.UpdateLifestyleData)
		analytics.DELETE("/lifestyle-data/:id", h.DeleteLifestyleData)

		// Wearable Data
		analytics.POST("/wearable-data", h.AddWearableData)
		analytics.GET("/wearable-data", h.GetWearableData)
		analytics.GET("/wearable-data/:id", h.GetWearableDataById)
		analytics.PUT("/wearable-data", h.UpdateWearableData)
		analytics.DELETE("/wearable-data/:id", h.DeleteWearableData)

		// Health Recommendations
		analytics.POST("/health-recommendations", h.GenerateHealthRecommendations)
		analytics.GET("/health-recommendations/:id", h.GetHealthRecommendationsById)

		// Health Monitoring
		analytics.GET("/health-monitoring/realtime", h.GetRealtimeHealthMonitoring)
		analytics.GET("/health-monitoring/daily-summary", h.GetDailyHealthSummary)
		analytics.GET("/health-monitoring/weekly-summary", h.GetWeeklyHealthSummary)
	}

	swaggerURL := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, swaggerURL))

	return r
}
