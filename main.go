package main

import (
	"FUS/utils"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env:", err)
	}

	PORT := os.Getenv("PORT")
	fmt.Println("PORT value:", PORT)

	utils.SetupDB()

	app := gin.Default()

	if err := app.Run(":" + PORT); err != nil {
		log.Fatal(err)
	}
	fmt.Println("hello from file upload system!")
}
