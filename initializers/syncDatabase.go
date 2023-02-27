package initializers

import "github.com/dzeleniak/jwt-api/models"

func SyncDatabase() {
	// Create user table on load
	DB.AutoMigrate(&models.User{})
}
