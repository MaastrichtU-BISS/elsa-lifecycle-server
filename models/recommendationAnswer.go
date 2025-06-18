package models

import (
	"github.com/google/uuid"
)

type RecommendationAnswer struct {
	ID               uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Form             string         `json:"form"`
	File             string         `json:"file" gorm:"type:varchar(500)"` // Path or URL to file
	RecommendationID uint           `json:"recommendationId"`
	Recommendation   Recommendation `gorm:"foreignKey:RecommendationID"`
	UserID           uuid.UUID      `json:"userId"`
	User             User           `gorm:"foreignKey:UserID"`
}
