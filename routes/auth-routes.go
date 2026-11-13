package routes

import (
	"FUS/controllers"
	"FUS/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	authRoutes := app.Group("/auth")

	authRoutes.POST("/login", controllers.HandleLogin)
	authRoutes.POST("/register", controllers.HandleRegister)
	authRoutes.POST("/logout", middleware.AuthRequired, controllers.HandleLogout)
	authRoutes.GET("/me", middleware.AuthRequired, controllers.GetUser)
}
