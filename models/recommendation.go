package models

type Recommendation struct {
	ID               uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	QuestionnaireID  uint          `json:"questionnaireId"`
	Questionnaire    Questionnaire `json:"questionnaire" gorm:"foreignKey:QuestionnaireID"`
	ToolID           uint          `json:"toolId"`
	Tool             Tool          `gorm:"foreignKey:ToolID"`
	BinaryEvaluation uint          `json:"binaryEvaluation"`
}
