package models

type Reflection struct {
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title          string `json:"title" gorm:"type:varchar(255);not null"`
	Description    string `json:"description"`
	Considerations string `json:"considerations" gorm:"type:text"`
	Form           string `json:"form" gorm:"type:text;not null"`
	PhaseID        uint   `json:"phaseId"`
	Phase          Phase  `gorm:"foreignKey:PhaseID"` // Relationship to the phase
}
