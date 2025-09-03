package main

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
	General      string `json:"General"`
	Introduction string `json:"Introduction"`
}

type PhaseSeed struct {
	Number      int    `json:"Number"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	LifecycleID int    `json:"LifecycleID"`
}

type ToolSeed struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	URL         string `json:"URL"`
	Cover       string `json:"Cover"`
	Tags        string `json:"Tags"`
	Type        string `json:"Type"`
	FormFile    string `json:"FormFile"`
	Form        string `json:"-"`
}

func main() {
	database.ConnectDB()
	db := database.DB

	// Users
	var users []UserSeed
	readSeed("database/seeds/users.json", &users)
	for _, u := range users {
		uuidVal, err := uuid.Parse(u.ID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid UUID for user: %s\n", u.ID)
			continue
		}
		db.Create(&models.User{
			ID:           uuidVal,
			Email:        u.Email,
			PasswordHash: u.PasswordHash,
		})
	}

	// Lifecycles
	var lifecycles []LifecycleSeed
	readSeed("database/seeds/lifecycles.json", &lifecycles)
	for _, l := range lifecycles {
		db.Create(&models.Lifecycle{
			Title:        l.Title,
			Description:  l.Description,
			General:      l.General,
			Introduction: l.Introduction,
		})
	}

	// Phases
	var phases []PhaseSeed
	readSeed("database/seeds/phases.json", &phases)
	for _, p := range phases {
		db.Create(&models.Phase{
			Number:      uint(p.Number),
			Title:       p.Title,
			Description: p.Description,
			LifecycleID: uint(p.LifecycleID),
		})
	}

	// Tools (with JSON-LD form)
	var tools []ToolSeed
	readSeed("database/seeds/tools.json", &tools)
	for i, t := range tools {
		if t.FormFile != "" {
			formData, err := os.ReadFile(filepath.Join("database/seeds", t.FormFile))
			if err == nil {
				tools[i].Form = string(formData)
			}
		}
		db.Create(&models.Tool{
			Title:       t.Title,
			Description: t.Description,
			URL:         t.URL,
			Cover:       t.Cover,
			Tags:        t.Tags,
			Type:        t.Type,
			Form:        tools[i].Form,
		})
	}

	// Reflections (with JSON-LD form)
	var reflections []struct {
		Title       string `json:"Title"`
		Description string `json:"Description"`
		FormFile    string `json:"FormFile"`
		Form        string `json:"-"`
		PhaseID     uint   `json:"PhaseID"`
	}
	readSeed("database/seeds/reflections.json", &reflections)
	for i, r := range reflections {
		if r.FormFile != "" {
			formData, err := os.ReadFile(filepath.Join("database/seeds", r.FormFile))
			if err == nil {
				reflections[i].Form = string(formData)
			}
		}
		db.Create(&models.Reflection{
			Title:       r.Title,
			Description: r.Description,
			Form:        reflections[i].Form,
			PhaseID:     r.PhaseID,
		})
	}

	// ReflectionAnswers
	var reflectionAnswers []struct {
		Form             string `json:"Form"`
		BinaryEvaluation uint   `json:"BinaryEvaluation"`
		ReflectionID     uint   `json:"ReflectionID"`
		UserID           string `json:"UserID"`
	}
	readSeed("database/seeds/reflection_answers.json", &reflectionAnswers)
	for _, ra := range reflectionAnswers {
		userID, err := uuid.Parse(ra.UserID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid UUID for reflection answer: %s\n", ra.UserID)
			continue
		}
		db.Create(&models.ReflectionAnswer{
			Form:             ra.Form,
			BinaryEvaluation: ra.BinaryEvaluation,
			ReflectionID:     ra.ReflectionID,
			UserID:           userID,
		})
	}

	// Journals (with JSON-LD form)
	var journals []struct {
		Title       string `json:"Title"`
		Description string `json:"Description"`
		FormFile    string `json:"FormFile"`
		Form        string `json:"-"`
		PhaseID     uint   `json:"PhaseID"`
	}
	readSeed("database/seeds/journals.json", &journals)
	for i, j := range journals {
		if j.FormFile != "" {
			formData, err := os.ReadFile(filepath.Join("database/seeds", j.FormFile))
			if err == nil {
				journals[i].Form = string(formData)
			}
		}
		db.Create(&models.Journal{
			Title:       j.Title,
			Description: j.Description,
			Form:        journals[i].Form,
			PhaseID:     j.PhaseID,
		})
	}

	// JournalAnswers
	var journalAnswers []struct {
		Form      string `json:"Form"`
		JournalID uint   `json:"JournalID"`
		UserID    string `json:"UserID"`
	}
	readSeed("database/seeds/journal_answers.json", &journalAnswers)
	for _, ja := range journalAnswers {
		userID, err := uuid.Parse(ja.UserID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid UUID for journal answer: %s\n", ja.UserID)
			continue
		}
		db.Create(&models.JournalAnswer{
			Form:      ja.Form,
			JournalID: ja.JournalID,
			UserID:    userID,
		})
	}

	// Recommendations
	var recommendations []struct {
		ReflectionID     uint `json:"ReflectionID"`
		ToolID           uint `json:"ToolID"`
		BinaryEvaluation uint `json:"BinaryEvaluation"`
	}
	readSeed("database/seeds/recommendations.json", &recommendations)
	for _, rec := range recommendations {
		db.Create(&models.Recommendation{
			ReflectionID:     rec.ReflectionID,
			ToolID:           rec.ToolID,
			BinaryEvaluation: rec.BinaryEvaluation,
		})
	}

	// RecommendationAnswers
	var recommendationAnswers []struct {
		Form             string `json:"Form"`
		File             string `json:"File"`
		RecommendationID uint   `json:"RecommendationID"`
		UserID           string `json:"UserID"`
	}
	readSeed("database/seeds/recommendation_answers.json", &recommendationAnswers)
	for _, ra := range recommendationAnswers {
		userID, err := uuid.Parse(ra.UserID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid UUID for recommendation answer: %s\n", ra.UserID)
			continue
		}
		db.Create(&models.RecommendationAnswer{
			Form:             ra.Form,
			File:             ra.File,
			RecommendationID: ra.RecommendationID,
			UserID:           userID,
		})
	}

	fmt.Println("Seeding complete.")
}

func readSeed(path string, v interface{}) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
		os.Exit(1)
	}
	if err := json.Unmarshal(data, v); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", path, err)
		os.Exit(1)
	}
}
