package main

import (
	"FUS/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	utils.SetupDB()

	app := gin.Default()

	app.Run(fmt.Sprintf("%v", PORT))
	fmt.Println("hello from file upload system!")
}
