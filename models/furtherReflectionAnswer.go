package models

import "time"
import "github.com/google/uuid"

type FurtherReflectionAnswer struct {
ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
ReflectionID uint       `json:"reflectionId"`
Reflection   Reflection `gorm:"foreignKey:ReflectionID"`
UserID       uuid.UUID  `json:"userId"`
User         User       `gorm:"foreignKey:UserID"`
Form         string     `json:"form" gorm:"type:text"`
UpdatedAt    time.Time  `json:"updatedAt"`
}
