package models

import "time"

type FurtherReflectionAnswer struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	ReflectionID uint       `json:"reflectionId"`
	Reflection   Reflection `gorm:"foreignKey:ReflectionID"`
	JournalID    uint       `json:"journalId" gorm:"index:idx_further_reflection_user_updated"`
	Journal      Journal    `gorm:"foreignKey:JournalID"`
	Form         string     `json:"form" gorm:"type:text"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"index:idx_further_reflection_user_updated,sort:desc"`
}
