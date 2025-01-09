package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `gorm:"type:varchar(100);unique;not null"`
	Username string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}
