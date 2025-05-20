package models

type Reflection struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title        string `json:"title"`
	Description string `json:"description"`
	Form        string `json:"form"`
	PhaseID     uint   `json:"phaseId"`
	Phase       Phase  `gorm:"foreignKey:PhaseID"` // Relationship to the phase
	Answers     []ReflectionAnswer
}
