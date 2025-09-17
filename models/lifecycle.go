package models

type Lifecycle struct {
	ID           uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title        string  `json:"title" gorm:"type:varchar(255);not null"`
	Description  string  `json:"description" gorm:"type:text"`
	Welcome      string  `json:"welcome" gorm:"type:text"`
	Introduction string  `json:"introduction" gorm:"type:text"`
	Journal      string  `json:"journal" gorm:"type:text"`
	Phases       []Phase // Relationship to the phases
}
