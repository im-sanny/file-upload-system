package controllers

import (
	"FUS/models"
	"FUS/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "some fields are missing"})
		return
	}

	var dbUser models.User
	if err := utils.DB.Where("username = ?", user.Username).First(dbUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("user with %s username not found!", user.Username)})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	accessToken, _ := utils.GenerateToken(dbUser)
	c.SetCookie(
		"accessToken", accessToken,
		60*60*24*7,
		"/", "localhost",
		false, false,
	)

	c.JSON(http.StatusOK, gin.H{
		"message":     "logged in successfully",
		"accessToken": accessToken,
	})

}

func HandleRegister(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "some fields are missing"})
		return
	}

	var existingUser models.User
	err := utils.DB.Where("username = ? or email = ?", user.Username, user.Email).First(existingUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or email already exists!"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password!"})
		return
	}

	user.Password = string(hashedPassword)
	if err := utils.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user!"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func HandleLogout(c *gin.Context) {

}

func GetUser(c *gin.Context) {

}
