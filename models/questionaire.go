package models

type Questionnaire struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Form        int    `json:"form"`
}
