package controllers

import "testing"

func TestIsInRecommendationAnswerUploadDir(t *testing.T) {
	inside := []string{
		"uploads/recommendation_answers/0b1e/report.pdf",
		"./uploads/recommendation_answers/legacy.pdf",
	}
	for _, path := range inside {
		if !isInRecommendationAnswerUploadDir(path) {
			t.Fatalf("%q should be inside the upload dir", path)
		}
	}

	outside := []string{
		"",
		"uploads/recommendation_answers",
		"uploads/recommendation_answers/../../main.go",
		"uploads/cover.png",
		"/etc/passwd",
		"database/db/elsa.db",
	}
	for _, path := range outside {
		if isInRecommendationAnswerUploadDir(path) {
			t.Fatalf("%q should be outside the upload dir", path)
		}
	}
}
