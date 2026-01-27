package models

type Journal struct {
	ID           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title        string     `json:"title" gorm:"type:varchar(255)"`
	Form         string     `json:"form"`
	ReflectionID uint       `json:"reflectionId"`
	Reflection   Reflection `gorm:"foreignKey:ReflectionID"`
}
