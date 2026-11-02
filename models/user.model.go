package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"not nul;unique"`
	Email    string `json:"email" gorm:"not nul;unique"`
	Password string `json:"password" gorm:"not nul"`
}
