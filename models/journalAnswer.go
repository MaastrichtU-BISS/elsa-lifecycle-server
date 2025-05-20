package models

type JournalAnswer struct {
	ID               uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Form             string        `json:"form"`
	JournalID  uint          `json:"journalId"`
	Journal    Journal `gorm:"foreignKey:JournalID"`
}
