package api

import (
	"api/api/handler"
	// "api/api/middleware"
	// "log"

	// "github.com/casbin/casbin/v2"
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

	// ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	// if err != nil {
	// 	panic(err)
	// }

	// err = ca.LoadPolicy()
	// if err != nil {
	// 	log.Fatal("casbin error load policy: ", err)
	// 	panic(err)
	// }

	auth := r.Group("/auth")
	// auth.Use(middleware.NewAuth(ca))
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/validateToken", h.ValidateToken)
		auth.POST("/refreshToken", h.RefreshToken)
		auth.POST("/validateEmail", h.ValidateEmail)
		auth.POST("/verifyCode", h.VerificationCode)
	}

	user := r.Group("/users")
	// user.Use(middleware.NewAuth(ca))
	{
		user.GET("/profile/:user_id", h.GetProfileInfo)
		user.PUT("/updateProfile", h.UpdateProfile)
		user.DELETE("/deleteProfile/:user_id", h.DeleteProfile)
		user.POST("/change-password", h.ChangePassword)
	}

	swaggerURL := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, swaggerURL))

	return r
}
