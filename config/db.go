package config

import (
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgresql://postgres:postgres@postgres:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
