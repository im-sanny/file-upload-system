package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	fmt.Println("Connecting to db...")

	var err error
	DB, err = gorm.Open(sqlite.Open("db/sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database")
	}

	fmt.Println("Connected to db successfully")
}
