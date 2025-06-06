package models

type RecommendationAnswer struct {
	ID               uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Form             string     `json:"form"`
	File             string     `json:"cover" gorm:"type:varchar(500)"` // Path or URL to file
	RecommendationID uint       `json:"recommendationId"`
	Recommendation   Reflection `gorm:"foreignKey:RecommendationID"`
}
