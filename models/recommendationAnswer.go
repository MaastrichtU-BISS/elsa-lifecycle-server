package models

import (
	"time"
)

type RecommendationAnswer struct {
	ID               uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Form             string         `json:"form"`
	File             string         `json:"file" gorm:"type:varchar(500)"` // Path or URL to file
	CheckedDone      bool           `json:"checked_done" gorm:"default:false"`
	RecommendationID uint           `json:"recommendationId"`
	Recommendation   Recommendation `gorm:"foreignKey:RecommendationID"`
	JournalID        uint           `json:"journalId" gorm:"index:idx_recommendation_user_updated"`
	Journal          Journal        `gorm:"foreignKey:JournalID"`
	UpdatedAt        time.Time      `json:"updatedAt" gorm:"index:idx_recommendation_user_updated,sort:desc"`
}
