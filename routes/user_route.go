package routes

import (
	"github.com/gin-gonic/gin"
	"golang-web-server/config"
	"golang-web-server/controllers"
	"golang-web-server/repositories"
)

func SetupUserRoutes(r *gin.Engine) {
	userRepo := repositories.NewUserRepository(config.DB)
	userController := controllers.NewUserController(userRepo)

	user := r.Group("/api/user")
	{
		user.POST("/register", userController.RegisterNewUser)
		user.POST("/login", userController.Login)
		user.GET("/:id", userController.GetUser)
	}
}
