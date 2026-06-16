package seeder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"server/database"
	"server/models"

	"github.com/google/uuid"
)

type UserSeed struct {
	ID           string `json:"ID"`
	Email        string `json:"Email"`
	PasswordHash string `json:"PasswordHash"`
}

type LifecycleSeed struct {
	Title        string `json:"Title"`
	Description  string `json:"Description"`
	Introduction string `json:"Introduction"`
}

type PhaseSeed struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	LifecycleID int    `json:"LifecycleID"`
}

type ToolSeed struct {
	Title       string  `json:"Title"`
	Description string  `json:"Description"`
	URL         string  `json:"URL"`
	Cover       string  `json:"Cover"`
	Tags        *string `json:"Tags" gorm:"default:null"`
	Type        *string `json:"Type" gorm:"default:null"`
	FormFile    string  `json:"FormFile"`
	Form        string  `json:"-"`
	FileUpload  bool    `json:"FileUpload"`
}

func RequireTestEnvironment() error {
	if os.Getenv("APP_ENV") != "test" {
		return fmt.Errorf("refusing to reset database outside test environment")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		return fmt.Errorf("DB_PATH must be set when running test seed")
	}

	if filepath.Base(dbPath) == "elsa.db" {
		return fmt.Errorf("refusing to reset main database elsa.db")
	}

	return nil
}

func ResetAndSeedDatabase() error {
	database.ConnectDB()
	db := database.DB

	db.Migrator().DropTable(
		&models.User{},
		&models.Lifecycle{},
		&models.Phase{},
		&models.Tool{},
		&models.Reflection{},
		&models.ReflectionAnswer{},
		&models.FurtherReflectionAnswer{},
		&models.Recommendation{},
		&models.RecommendationAnswer{},
	)

	if err := db.AutoMigrate(
		&models.User{},
		&models.Lifecycle{},
		&models.Phase{},
		&models.Tool{},
		&models.Reflection{},
		&models.ReflectionAnswer{},
		&models.FurtherReflectionAnswer{},
		&models.Recommendation{},
		&models.RecommendationAnswer{},
	); err != nil {
		return err
	}

	var users []UserSeed
	if err := readSeed("database/seeds/users.json", &users); err != nil {
		return err
	}

	for _, u := range users {
		uuidVal, err := uuid.Parse(u.ID)
		if err != nil {
			return fmt.Errorf("invalid UUID for user %s: %w", u.ID, err)
		}

		if err := db.Create(&models.User{
			ID:           uuidVal,
			Email:        u.Email,
			PasswordHash: u.PasswordHash,
		}).Error; err != nil {
			return err
		}
	}

	var lifecycles []LifecycleSeed
	if err := readSeed("database/seeds/lifecycles.json", &lifecycles); err != nil {
		return err
	}

	for _, l := range lifecycles {
		if err := db.Create(&models.Lifecycle{
			Title:        l.Title,
			Description:  l.Description,
			Introduction: l.Introduction,
		}).Error; err != nil {
			return err
		}
	}

	var phases []PhaseSeed
	if err := readSeed("database/seeds/phases.json", &phases); err != nil {
		return err
	}

	for _, p := range phases {
		if err := db.Create(&models.Phase{
			Title:       p.Title,
			Description: p.Description,
			LifecycleID: uint(p.LifecycleID),
		}).Error; err != nil {
			return err
		}
	}

	var tools []ToolSeed
	if err := readSeed("database/seeds/tools.json", &tools); err != nil {
		return err
	}

	for i, t := range tools {
		if t.FormFile != "" {
			formData, err := os.ReadFile(filepath.Join("database/seeds", t.FormFile))
			if err != nil {
				return err
			}
			tools[i].Form = string(formData)
		}

		if err := db.Create(&models.Tool{
			Title:       t.Title,
			Description: t.Description,
			URL:         t.URL,
			Cover:       t.Cover,
			Tags:        t.Tags,
			Type:        t.Type,
			Form:        tools[i].Form,
			FileUpload:  t.FileUpload,
		}).Error; err != nil {
			return err
		}
	}

	var reflections []struct {
		FormFile                  string `json:"FormFile"`
		Form                      string `json:"Form"`
		FurtherReflectionFormFile string `json:"FurtherReflectionFormFile"`
		FurtherReflectionForm     string `json:"-"`
		Description               string `json:"Description"`
		Title                     string `json:"Title"`
		Considerations            string `json:"Considerations"`
		PhaseID                   uint   `json:"PhaseID"`
	}

	if err := readSeed("database/seeds/reflections.json", &reflections); err != nil {
		return err
	}

	for i, r := range reflections {
		if r.FormFile != "" {
			formData, err := os.ReadFile(filepath.Join("database/seeds", r.FormFile))
			if err != nil {
				return err
			}
			reflections[i].Form = string(formData)
		}

		if r.FurtherReflectionFormFile != "" {
			formData, err := os.ReadFile(filepath.Join("database/seeds", r.FurtherReflectionFormFile))
			if err != nil {
				return err
			}
			reflections[i].FurtherReflectionForm = string(formData)
		}

		if err := db.Create(&models.Reflection{
			Form:                  reflections[i].Form,
			FurtherReflectionForm: reflections[i].FurtherReflectionForm,
			Description:           r.Description,
			Title:                 r.Title,
			Considerations:        r.Considerations,
			PhaseID:               r.PhaseID,
		}).Error; err != nil {
			return err
		}
	}

	var reflectionAnswers []struct {
		Form         string `json:"Form"`
		ReflectionID uint   `json:"ReflectionID"`
		UserID       string `json:"UserID"`
	}

	if err := readSeed("database/seeds/reflection_answers.json", &reflectionAnswers); err != nil {
		return err
	}

	for _, ra := range reflectionAnswers {
		userID, err := uuid.Parse(ra.UserID)
		if err != nil {
			return fmt.Errorf("invalid UUID for reflection answer %s: %w", ra.UserID, err)
		}

		if err := db.Create(&models.ReflectionAnswer{
			Form:         ra.Form,
			ReflectionID: ra.ReflectionID,
			UserID:       userID,
		}).Error; err != nil {
			return err
		}
	}

	var furtherReflectionAnswers []struct {
		Form         string `json:"Form"`
		ReflectionID uint   `json:"ReflectionID"`
		UserID       string `json:"UserID"`
	}

	if err := readSeed("database/seeds/further_reflection_answers.json", &furtherReflectionAnswers); err != nil {
		return err
	}

	for _, fra := range furtherReflectionAnswers {
		userID, err := uuid.Parse(fra.UserID)
		if err != nil {
			return fmt.Errorf("invalid UUID for further reflection answer %s: %w", fra.UserID, err)
		}

		if err := db.Create(&models.FurtherReflectionAnswer{
			Form:         fra.Form,
			ReflectionID: fra.ReflectionID,
			UserID:       userID,
		}).Error; err != nil {
			return err
		}
	}

	var recommendations []struct {
		ReflectionID     uint `json:"ReflectionID"`
		ToolID           uint `json:"ToolID"`
		BinaryEvaluation uint `json:"BinaryEvaluation"`
	}

	if err := readSeed("database/seeds/recommendations.json", &recommendations); err != nil {
		return err
	}

	for _, rec := range recommendations {
		if err := db.Create(&models.Recommendation{
			ReflectionID: rec.ReflectionID,
			ToolID:       rec.ToolID,
		}).Error; err != nil {
			return err
		}
	}

	var recommendationAnswers []struct {
		Form             string `json:"Form"`
		File             string `json:"File"`
		RecommendationID uint   `json:"RecommendationID"`
		UserID           string `json:"UserID"`
	}

	if err := readSeed("database/seeds/recommendation_answers.json", &recommendationAnswers); err != nil {
		return err
	}

	for _, ra := range recommendationAnswers {
		userID, err := uuid.Parse(ra.UserID)
		if err != nil {
			return fmt.Errorf("invalid UUID for recommendation answer %s: %w", ra.UserID, err)
		}

		if err := db.Create(&models.RecommendationAnswer{
			Form:             ra.Form,
			File:             ra.File,
			RecommendationID: ra.RecommendationID,
			UserID:           userID,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

func readSeed(path string, v interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading %s: %w", path, err)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("error parsing %s: %w", path, err)
	}

	return nil
}
