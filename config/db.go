package config

import (
	"awesomeNotes/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase() error {
	dbUri := "host=localhost port=5432 user=login password=pass dbname=notes2"
	var err error
	DB, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Note{})
	return nil
}
