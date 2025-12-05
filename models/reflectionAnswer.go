package models

import "github.com/google/uuid"

type ReflectionAnswer struct {
	ID                 uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	ReflectionID       uint       `json:"reflectionId"`
	Reflection         Reflection `gorm:"foreignKey:ReflectionID"`
	UserID             uuid.UUID  `json:"userId"`
	User               User       `gorm:"foreignKey:UserID"`
	AnswerText         string     `json:"answerText" gorm:"type:text"`
	GetRecommendations bool       `json:"getRecommendations"`
}
