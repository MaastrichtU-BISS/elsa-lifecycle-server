package models

type Questionnaire struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Form        string `json:"form"`
	FormName    string `json:"formName"`
}
