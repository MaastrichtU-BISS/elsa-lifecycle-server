package database

import (
	"log"
	"os"
	"server/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// default connection string, matching the Postgres service in docker-compose.yml
const defaultDatabaseURL = "postgres://elsa:elsa@localhost:5432/elsa?sslmode=disable"

// DatabaseURL returns the connection string from DATABASE_URL, or the local default
func DatabaseURL() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}
	return defaultDatabaseURL
}

func ConnectDB() {
	var err error

	DB, err = gorm.Open(postgres.Open(DatabaseURL()), &gorm.Config{
		// map constraint violations to gorm.ErrForeignKeyViolated / gorm.ErrDuplicatedKey
		TranslateError: true,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get database handle: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := DB.AutoMigrate(&models.Lifecycle{},
		&models.Reflection{},
		&models.Journal{},
		&models.ReflectionAnswer{},
		&models.FurtherReflectionAnswer{},
		&models.Tool{},
		&models.Recommendation{},
		&models.RecommendationAnswer{},
		&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
