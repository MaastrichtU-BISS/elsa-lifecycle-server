package models

type Tool struct {
	ID          uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string  `json:"title" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:text"`
	URL         string  `json:"url" gorm:"type:varchar(500)"`
	Form        string  `json:"form"`
	FileUpload  bool    `json:"file_upload" gorm:"default:false"`
	Cover       string  `json:"cover" gorm:"type:varchar(500)"` // Path or URL to image
	Tags        *string `json:"tags" gorm:"default:null"`       // Comma-separated tags
	Type        *string `json:"type" gorm:"default:null"`       // Comma-separated tags
}
