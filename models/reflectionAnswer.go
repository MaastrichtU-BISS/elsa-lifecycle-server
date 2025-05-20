package models

type ReflectionAnswer struct {
	ID               uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Form             string        `json:"form"`
	BinaryEvaluation uint          `json:"binaryEvaluation"`
	ReflectionID  uint          `json:"reflectionId"`
	Reflection    Reflection `gorm:"foreignKey:ReflectionID"`
}
