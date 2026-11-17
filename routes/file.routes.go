package routes

import (
	"FUS/controllers"
	"FUS/middleware"

	"github.com/gin-gonic/gin"
)

func FileRoutes(router gin.RouterGroup) {
	file_router := router.Group("/file")

	file_router.POST("/upload", middleware.AuthRequired, controllers.UploadFile)
	file_router.GET("/", middleware.AuthRequired, controllers.ListFiles)
	file_router.GET("/:filename", middleware.AuthRequired, controllers.GetFiles)
	file_router.DELETE("/:filename", middleware.AuthRequired, controllers.DeleteFiles)
}
