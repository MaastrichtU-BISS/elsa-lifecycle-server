package database

import (
	"fmt"
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

	// Only run seeders if SEED environment variable is set to TRUE (case-insensitive)
	if os.Getenv("SEED") == "TRUE" || os.Getenv("SEED") == "true" {
		seeders := []Seeder{
			LifecycleSeeder{},
			PhaseSeeder{},
			ReflectionSeeder{},
			ReflectionAnswerSeeder{},
			JournalSeeder{},
			JournalAnswerSeeder{},
			ToolSeeder{},
			RecommendationSeeder{},
			RecommendationAnswerSeeder{},
			UserSeeder{},
		}

		for _, seeder := range seeders {
			if err := seeder.Clear(DB); err != nil {
				fmt.Printf("Clear failed: %v", err)
			}

			if err := seeder.Seed(DB); err != nil {
				fmt.Printf("Seeding failed: %v", err)
			}
		}
	}
}
