package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	Email string	`gorm:"unique"`
	Username string	`gorm:"unique"`
	Password string
  }