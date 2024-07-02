package initializers

import "github.com/anojaryal/JWT-Authentication/models"


func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}