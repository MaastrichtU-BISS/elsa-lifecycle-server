package models

// Reflection is one section of a lifecycle's journal (e.g. "Problem Definition")
type Reflection struct {
	ID                    uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title                 string    `json:"title" gorm:"type:varchar(255);not null"`
	Context               string    `json:"context" gorm:"type:text"` // Introductory paragraph shown above the question
	Description           string    `json:"description"`              // The question the user answers
	Considerations        string    `json:"considerations" gorm:"type:text"`
	Form                  string    `json:"form" gorm:"type:text;not null"`
	FurtherReflectionForm string    `json:"furtherReflectionForm" gorm:"type:text"`
	LifecycleID           uint      `json:"lifecycleId"`
	Lifecycle             Lifecycle `json:"-" gorm:"foreignKey:LifecycleID"` // Relationship to the lifecycle
}
