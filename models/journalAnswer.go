package models

import "github.com/google/uuid"

type JournalAnswer struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Form      string    `json:"form"`
	JournalID uint      `json:"journalId"`
	Journal   Journal   `gorm:"foreignKey:JournalID"`
	UserID    uuid.UUID `json:"userId"`
	User      User      `gorm:"foreignKey:UserID"`
}
