package models

type Lifecycle struct {
	ID          uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string  `json:"title" gorm:"type:varchar(255);not null"` // Title of the lifecycle
	Description string  `json:"description" gorm:"type:text"`            // Description of the lifecycle
	Phases      []Phase // Relationship to the phases
}
