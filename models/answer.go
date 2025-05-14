package models

type Answer struct {
	ID               uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Form             string        `json:"form"`
	BinaryEvaluation uint          `json:"binaryEvaluation"`
	QuestionnaireID  uint          `json:"questionnaireId"`
	Questionnaire    Questionnaire `gorm:"foreignKey:QuestionnaireID"`
}
