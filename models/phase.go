package models

type Phase struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Number        uint      `json:"number" gorm:"not null"`                  // Phase number
	Title         string    `json:"title" gorm:"type:varchar(255);not null"` // Title of the phase
	Description   string    `json:"description" gorm:"type:text"`            // Description of the phase
	LifecycleID   uint      `json:"lifecycleId"`                             // Foreign key to the lifecycle
	Lifecycle     Lifecycle `gorm:"foreignKey:LifecycleID"`                  // Relationship to the lifecycle
	Reflection    *Reflection
	Journal       *Journal
}
