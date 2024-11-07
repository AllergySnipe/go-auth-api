package initializer

import "github.com/AllergySnipe/go-auth-api/models"

func MigrateDb() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.RevokedToken{})
}
