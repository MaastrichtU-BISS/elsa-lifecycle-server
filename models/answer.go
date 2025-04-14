package models

type Answer struct {
	ID              uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Form            string `json:"form"`
	QuestionnaireID uint
	Questionnaire   Questionnaire `gorm:"foreignKey:QuestionnaireID"`
}
