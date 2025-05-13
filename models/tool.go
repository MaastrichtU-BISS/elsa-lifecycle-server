package models

type Tool struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text"`
	URL         string `json:"url" gorm:"type:varchar(500);not null"`
	Cover       string `json:"cover" gorm:"type:varchar(500)"` // Path or URL to image
}
