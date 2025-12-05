package models

type Phase struct {
	ID          uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string       `json:"title" gorm:"type:varchar(255);not null"` // Title of the phase
	Description string       `json:"description" gorm:"type:text"`            // Description of the phase
	LifecycleID uint         `json:"lifecycleId"`                             // Foreign key to the lifecycle
	Lifecycle   Lifecycle    `gorm:"foreignKey:LifecycleID"`                  // Relationship to the lifecycle
	Reflections []Reflection `gorm:"foreignKey:PhaseID"`                      // Relationship to reflections
	Journal     *Journal
}
