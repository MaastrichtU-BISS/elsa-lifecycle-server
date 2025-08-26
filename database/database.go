package database

import (
	"os"
	"log"

	"server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// default path (absolute path used in the Docker image)
	defaultPath := "/app/database/db/elsa.db"
	// allow override from environment variable DB_PATH
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultPath
	}

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database (%s): %v", dbPath, err)
	}

	DB.AutoMigrate(&models.Questionnaire{}, models.Answer{})
}
