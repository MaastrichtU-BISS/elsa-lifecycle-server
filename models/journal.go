package models

import (
	"time"

	"github.com/google/uuid"
)

type Journal struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	UserID      uuid.UUID `json:"userId"`
	User        User      `gorm:"foreignKey:UserID"` // Relationship to the user
	LifecycleID uint      `json:"lifecycleId"`
	Lifecycle   Lifecycle `gorm:"foreignKey:LifecycleID"` // Relationship to the lifecycle
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
