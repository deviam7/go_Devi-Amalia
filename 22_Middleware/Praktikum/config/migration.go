package config

import "DEVI-AMALIA/models"

func InitialMigration() {
	var DB = GetDB()

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
