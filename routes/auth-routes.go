package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(app *gin.Engine) {
	authRoutes := app.Group("/auth")
}
