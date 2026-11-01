package routes

import (
	"FUS/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	authRoutes := app.Group("/auth")

	authRoutes.POST("/login", controllers.HandleLogin)
	authRoutes.POST("/register", controllers.HandleRegister)
	authRoutes.POST("/logout", controllers.HandleLogout)
	authRoutes.GET("/me", controllers.GetUser)
}
