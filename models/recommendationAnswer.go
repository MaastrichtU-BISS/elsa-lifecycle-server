package models

import "time"
import (
	"github.com/google/uuid"
)

type RecommendationAnswer struct {
ID               uint           `json:"id" gorm:"primaryKey;autoIncrement"`
Form             string         `json:"form"`
File             string         `json:"file" gorm:"type:varchar(500)"` // Path or URL to file
CheckedDone      bool           `json:"checked_done" gorm:"default:false"`
RecommendationID uint           `json:"recommendationId"`
Recommendation   Recommendation `gorm:"foreignKey:RecommendationID"`
UserID           uuid.UUID      `json:"userId" gorm:"index:idx_recommendation_user_updated"`
User             User           `gorm:"foreignKey:UserID"`
UpdatedAt        time.Time      `json:"updatedAt" gorm:"index:idx_recommendation_user_updated,sort:desc"`
}
