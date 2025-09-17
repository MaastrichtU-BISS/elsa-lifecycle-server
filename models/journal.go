package models

type Journal struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Form    string `json:"form"`
	PhaseID uint   `json:"phaseId"`
	Phase   Phase  `gorm:"foreignKey:PhaseID"`
}
