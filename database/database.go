package database

import (
	"fmt"
	"server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database/db/elsa.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Lifecycle{},
		&models.Phase{},
		&models.Reflection{},
		&models.Journal{},
		&models.ReflectionAnswer{},
		&models.JournalAnswer{},
		&models.Tool{},
		&models.Recommendation{},
		&models.RecommendationAnswer{})

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
