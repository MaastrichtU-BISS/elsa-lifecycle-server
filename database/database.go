package database

import (
	"log"
	"os"
	"server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	// default path (relative path used both local and Docker environments)
	defaultPath := "./database/db/elsa.db"
	// allow override from environment variable DB_PATH
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultPath
	}

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database (%s): %v", dbPath, err)
	}

	DB.AutoMigrate(&models.Lifecycle{},
		&models.Phase{},
		&models.Reflection{},
		&models.Journal{},
		&models.ReflectionAnswer{},
		&models.JournalAnswer{},
		&models.Tool{},
		&models.Recommendation{},
		&models.RecommendationAnswer{},
		&models.User{})
}
