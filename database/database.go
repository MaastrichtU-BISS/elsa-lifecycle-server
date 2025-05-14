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

	DB.AutoMigrate(&models.Questionnaire{}, models.Answer{}, models.Tool{}, models.Recommendation{})

	seeders := []Seeder{
		ToolSeeder{},
		QuestionnaireSeeder{},
		RecommendationSeeder{},
		AnswerSeeder{},
		// Add other seeders here
	}

	for _, seeder := range seeders {
		if err := seeder.Clear(DB); err != nil {
			fmt.Print("Clear failed: %v", err)
		}

		if err := seeder.Seed(DB); err != nil {
			fmt.Print("Seeding failed: %v", err)
		}
	}
}
