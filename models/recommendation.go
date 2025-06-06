package models

type Recommendation struct {
	ID                    uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	ReflectionID          uint       `json:"reflectionId"`
	Reflection            Reflection `json:"reflection" gorm:"foreignKey:ReflectionID"`
	ToolID                uint       `json:"toolId"`
	Tool                  Tool       `gorm:"foreignKey:ToolID"`
	BinaryEvaluation      uint       `json:"binaryEvaluation"`
	RecommendationAnswers []RecommendationAnswer
}
