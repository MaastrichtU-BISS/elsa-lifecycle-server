package models

import "github.com/google/uuid"

type ReflectionAnswer struct {
	ID               uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Form             string     `json:"form"`
	BinaryEvaluation uint       `json:"binaryEvaluation"`
	ReflectionID     uint       `json:"reflectionId"`
	Reflection       Reflection `gorm:"foreignKey:ReflectionID"`
	UserID           uuid.UUID  `json:"userId"`
	User             User       `gorm:"foreignKey:UserID"`
}
