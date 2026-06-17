package models

import (
	"time"
)

type ReflectionAnswer struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	ReflectionID uint       `json:"reflectionId"`
	Reflection   Reflection `gorm:"foreignKey:ReflectionID"`
	JournalID    uint       `json:"journalId"`
	Journal      Journal    `gorm:"foreignKey:JournalID"`
	Form         string     `json:"form" gorm:"type:text"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}
