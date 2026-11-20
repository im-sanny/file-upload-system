package controllers

import (
	"FUS/models"
	"FUS/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadDir = "./uploads"

func UploadFile(c *gin.Context) {
	userId, exists := c.Get("userId")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not get user id"})
		return
	}

	var user models.User
	if err := utils.DB.Find(&user, userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch user"})
		return
	}

	// retrieve the file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}

	// validate file size
	file_size_mb := file.Size / (1024 * 1024)
	if file_size_mb > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file too large! limit is 100MB"})
		return
	}

	filename := filepath.Base(file.Filename)
	useDir := filepath.Join(uploadDir, user.Username)
	if err := os.MkdirAll(useDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user directory"})
		return
	}

	dst := filepath.Join(uploadDir, user.Username, filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save the file"})
		return
	}

	fileModel := models.File{
		UserId:   user.ID,
		Path:     dst,
		FileName: filename,
		Size:     file.Size,
	}

	if err := utils.DB.Create(&fileModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file instance"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "file uploaded successfully",
		"url":     fmt.Sprintf("/file/%s", file.Filename),
	})
}

func ListFiles(c *gin.Context) {

}

func GetFiles(c *gin.Context) {

}

func DeleteFiles(c *gin.Context) {

}
