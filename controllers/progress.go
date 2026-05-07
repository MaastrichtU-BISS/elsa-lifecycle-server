package controllers

import (
	"net/http"
	"server/database"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
)

type LastUpdatedItem struct {
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	LifecycleID    uint      `json:"lifecycleId"`
	LifecycleTitle string    `json:"lifecycleTitle"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func pickLatest(current *LastUpdatedItem, candidate LastUpdatedItem) *LastUpdatedItem {
	if current == nil || candidate.UpdatedAt.After(current.UpdatedAt) {
		item := candidate
		return &item
	}
	return current
}

// GET /progress/last-updated - Get the last updated answer (reflection, further reflection, or recommendation) for the user
func GetLastUpdatedLifecycleItemForUser(c *gin.Context) {
	userId := c.GetString("user_id")
	db := database.DB
	var latest *LastUpdatedItem

	var reflectionAnswer models.ReflectionAnswer
	reflectionTx := db.
		Preload("Reflection").
		Preload("Reflection.Phase").
		Preload("Reflection.Phase.Lifecycle").
		Where("user_id = ?", userId).
		Order("updated_at DESC").
		Limit(1).
		Find(&reflectionAnswer)
	if reflectionTx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reflection answer"})
		return
	}
	if reflectionTx.RowsAffected > 0 {
		latest = pickLatest(latest, LastUpdatedItem{
			Type:           "reflection",
			Title:          reflectionAnswer.Reflection.Title,
			LifecycleID:    reflectionAnswer.Reflection.Phase.LifecycleID,
			LifecycleTitle: reflectionAnswer.Reflection.Phase.Lifecycle.Title,
			UpdatedAt:      reflectionAnswer.UpdatedAt,
		})
	}

	var furtherReflectionAnswer models.FurtherReflectionAnswer
	furtherTx := db.
		Preload("Reflection").
		Preload("Reflection.Phase").
		Preload("Reflection.Phase.Lifecycle").
		Where("user_id = ?", userId).
		Order("updated_at DESC").
		Limit(1).
		Find(&furtherReflectionAnswer)
	if furtherTx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch further reflection answer"})
		return
	}
	if furtherTx.RowsAffected > 0 {
		latest = pickLatest(latest, LastUpdatedItem{
			Type:           "further_reflection",
			Title:          furtherReflectionAnswer.Reflection.Title,
			LifecycleID:    furtherReflectionAnswer.Reflection.Phase.LifecycleID,
			LifecycleTitle: furtherReflectionAnswer.Reflection.Phase.Lifecycle.Title,
			UpdatedAt:      furtherReflectionAnswer.UpdatedAt,
		})
	}

	var recommendationAnswer models.RecommendationAnswer
	recommendationTx := db.
		Preload("Recommendation").
		Preload("Recommendation.Reflection").
		Preload("Recommendation.Reflection.Phase").
		Preload("Recommendation.Reflection.Phase.Lifecycle").
		Where("user_id = ?", userId).
		Order("updated_at DESC").
		Limit(1).
		Find(&recommendationAnswer)
	if recommendationTx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recommendation answer"})
		return
	}
	if recommendationTx.RowsAffected > 0 {
		latest = pickLatest(latest, LastUpdatedItem{
			Type:           "recommendation",
			Title:          recommendationAnswer.Recommendation.Reflection.Title,
			LifecycleID:    recommendationAnswer.Recommendation.Reflection.Phase.LifecycleID,
			LifecycleTitle: recommendationAnswer.Recommendation.Reflection.Phase.Lifecycle.Title,
			UpdatedAt:      recommendationAnswer.UpdatedAt,
		})
	}

	if latest == nil {
		c.JSON(http.StatusOK, nil)
		return
	}

	c.JSON(http.StatusOK, latest)
}
