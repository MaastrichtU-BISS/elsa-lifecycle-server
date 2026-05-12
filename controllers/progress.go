package controllers

import (
	"errors"
	"net/http"
	"server/database"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type LastUpdatedItem struct {
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	ToolID         uint      `json:"toolId,omitempty"`
	LifecycleID    uint      `json:"lifecycleId"`
	LifecycleTitle string    `json:"lifecycleTitle"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

const (
	ActivityReflection        = "reflection"
	ActivityFurtherReflection = "further_reflection"
	ActivityRecommendation    = "recommendation"
)

type latestLifecycleResults struct {
	reflection        *models.ReflectionAnswer
	furtherReflection *models.FurtherReflectionAnswer
	recommendation    *models.RecommendationAnswer
}

func fetchFirst[T any](query *gorm.DB, dest *T) error {
	return query.First(dest).Error
}

func preloadReflectionLifecycle(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Reflection").
		Preload("Reflection.Phase").
		Preload("Reflection.Phase.Lifecycle")
}

func latestByUser(db *gorm.DB, userID string) *gorm.DB {
	return db.
		Where("user_id = ?", userID).
		Order("updated_at DESC")
}

func preloadRecommendationLifecycle(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Recommendation").
		Preload("Recommendation.Reflection").
		Preload("Recommendation.Reflection.Phase").
		Preload("Recommendation.Reflection.Phase.Lifecycle")
}

func toLastUpdatedItemReflection(a models.ReflectionAnswer) LastUpdatedItem {
	return LastUpdatedItem{
		Type:           ActivityReflection,
		Title:          a.Reflection.Title,
		LifecycleID:    a.Reflection.Phase.LifecycleID,
		LifecycleTitle: a.Reflection.Phase.Lifecycle.Title,
		UpdatedAt:      a.UpdatedAt,
	}
}

func toLastUpdatedItemFurtherReflection(a models.FurtherReflectionAnswer) LastUpdatedItem {
	return LastUpdatedItem{
		Type:           ActivityFurtherReflection,
		Title:          a.Reflection.Title,
		LifecycleID:    a.Reflection.Phase.LifecycleID,
		LifecycleTitle: a.Reflection.Phase.Lifecycle.Title,
		UpdatedAt:      a.UpdatedAt,
	}
}

func toLastUpdatedItemRecommendation(a models.RecommendationAnswer) LastUpdatedItem {
	return LastUpdatedItem{
		Type:           ActivityRecommendation,
		Title:          a.Recommendation.Reflection.Title,
		ToolID:         a.Recommendation.ToolID,
		LifecycleID:    a.Recommendation.Reflection.Phase.LifecycleID,
		LifecycleTitle: a.Recommendation.Reflection.Phase.Lifecycle.Title,
		UpdatedAt:      a.UpdatedAt,
	}
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
	userID := c.GetString("user_id")

	results := latestLifecycleResults{}

	g, ctx := errgroup.WithContext(c.Request.Context())
	db := database.DB.WithContext(ctx)

	g.Go(func() error {
		var a models.ReflectionAnswer
		err := fetchFirst(latestByUser(preloadReflectionLifecycle(db), userID), &a)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		results.reflection = &a
		return nil
	})

	g.Go(func() error {
		var a models.FurtherReflectionAnswer
		err := fetchFirst(latestByUser(preloadReflectionLifecycle(db), userID), &a)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		results.furtherReflection = &a
		return nil
	})

	g.Go(func() error {
		var a models.RecommendationAnswer
		err := fetchFirst(latestByUser(preloadRecommendationLifecycle(db), userID), &a)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		results.recommendation = &a
		return nil
	})

	if err := g.Wait(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	var latest *LastUpdatedItem
	if results.reflection != nil {
		latest = pickLatest(latest, toLastUpdatedItemReflection(*results.reflection))
	}
	if results.furtherReflection != nil {
		latest = pickLatest(latest, toLastUpdatedItemFurtherReflection(*results.furtherReflection))
	}
	if results.recommendation != nil {
		latest = pickLatest(latest, toLastUpdatedItemRecommendation(*results.recommendation))
	}
	
	c.JSON(http.StatusOK, gin.H{"data": latest})
}
