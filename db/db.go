package db

import (
	"log"
	"school_council/config"
	"school_council/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.C.DbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.Student{}, &models.Grade{})
}
