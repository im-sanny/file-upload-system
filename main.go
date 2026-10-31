package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	app := gin.Default()

	app.Run(fmt.Sprintf("%v", PORT))
	fmt.Println("hello from file upload system!")
}
