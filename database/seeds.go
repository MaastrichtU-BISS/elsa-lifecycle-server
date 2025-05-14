package database

import (
	"server/models"

	"gorm.io/gorm"
)

type Seeder interface {
	Seed(db *gorm.DB) error
	Clear(db *gorm.DB) error
}

type ToolSeeder struct{}

func (s ToolSeeder) Seed(db *gorm.DB) error {
	tools := []models.Tool{
		{
			Title:       "Value Sensitive Design",
			Description: "A method for integrating stakeholder values into the design process.",
			URL:         "https://vsdesign.org/",
			Cover:       "",
			Tags:        "Values,Stakeholders",
			Type:        "Method",
		},
		{
			Title:       "Problem Framing Canvas",
			Description: "A visual tool to help teams critically reflect on how a problem is framed.",
			URL:         "https://realkm.com/wp-content/uploads/2023/05/Problem-Framing-Canvas-Handbook.pdf",
			Cover:       "",
			Tags:        "Problem,Framing",
			Type:        "Mapping, Tool, Canvas",
		},
		{
			Title:       "Stakeholder Mapping Tool",
			Description: "Tool for identifying and analyzing stakeholder relationships.",
			URL:         "https://www.stakeholdermap.com/",
			Cover:       "",
			Tags:        "Stakeholders",
			Type:        "Mapping, Tool",
		},
		{
			Title:       "UNESCO Ethics of AI Recommendation",
			Description: "Normative framework providing ethical guidelines for AI development and use.",
			URL:         "https://unesdoc.unesco.org/ark:/48223/pf0000381137",
			Cover:       "",
			Tags:        "Ethics",
			Type:        "Guidelines, Framework",
		},
		{
			Title:       "Data Ethics Canvas",
			Description: "Canvas for identifying and addressing ethical issues in data projects.",
			URL:         "https://theodi.org/article/data-ethics-canvas/",
			Cover:       "",
			Tags:        "Ethics",
			Type:        "Mapping, Canvas",
		},
	}
	return db.Create(&tools).Error
}

func (s ToolSeeder) Clear(db *gorm.DB) error {
	return db.Exec("DELETE FROM users").Error
}
