package seeder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"server/database"
	"server/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
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

type Options struct {
	// SkipUsers leaves out the demo users from users.json (use in production)
	SkipUsers bool
}

// HasUserData reports whether the database contains journals or users that did not come from
// the seed files, i.e. data that a reset would destroy
func HasUserData() (bool, error) {
	if database.DB == nil {
		database.ConnectDB()
	}
	db := database.DB

	var journals int64
	if err := db.Model(&models.Journal{}).Count(&journals).Error; err != nil {
		return false, err
	}
	if journals > 0 {
		return true, nil
	}

	var seedUsers []UserSeed
	if err := readSeed("database/seeds/users.json", &seedUsers); err != nil {
		return false, err
	}
	seedEmails := make([]string, 0, len(seedUsers))
	for _, u := range seedUsers {
		seedEmails = append(seedEmails, u.Email)
	}

	query := db.Model(&models.User{})
	if len(seedEmails) > 0 {
		query = query.Where("email NOT IN ?", seedEmails)
	}
	var users int64
	if err := query.Count(&users).Error; err != nil {
		return false, err
	}
	return users > 0, nil
}

func RequireTestEnvironment() error {
	if os.Getenv("APP_ENV") != "test" {
		return fmt.Errorf("refusing to reset database outside test environment")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL must be set when running test seed")
	}

	config, err := pgconn.ParseConfig(databaseURL)
	if err != nil {
		return fmt.Errorf("invalid DATABASE_URL: %w", err)
	}

	if config.Database == "elsa" {
		return fmt.Errorf("refusing to reset main database elsa")
	}

	return nil
}

func ResetAndSeedDatabase(opts Options) error {
	// reuse the server's connection when called from the reset endpoint
	if database.DB == nil {
		database.ConnectDB()
	}
	db := database.DB

	// phases were removed: reflections now belong directly to a lifecycle
	if err := db.Migrator().DropTable("phases"); err != nil {
		return err
	}

	if err := db.Migrator().DropTable(
		&models.Journal{},
		&models.User{},
		&models.Lifecycle{},
		&models.Tool{},
		&models.Reflection{},
		&models.ReflectionAnswer{},
		&models.FurtherReflectionAnswer{},
		&models.Recommendation{},
		&models.RecommendationAnswer{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&models.Journal{},
		&models.User{},
		&models.Lifecycle{},
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
	if !opts.SkipUsers {
		if err := readSeed("database/seeds/users.json", &users); err != nil {
			return err
		}
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
		Title                     string `json:"Title"`
		Context                   string `json:"Context"`
		Description               string `json:"Description"`
		Considerations            string `json:"Considerations"`
		LifecycleID               uint   `json:"LifecycleID"`
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
			Title:                 r.Title,
			Context:               r.Context,
			Description:           r.Description,
			Considerations:        r.Considerations,
			LifecycleID:           r.LifecycleID,
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
