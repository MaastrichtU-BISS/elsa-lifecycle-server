package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /tools/:id/tools - Fetch all tools
func GetTools(c *gin.Context) {
	var tools []models.Tool
	database.DB.Find(&tools)
	c.JSON(http.StatusOK, tools)
}

// GET /tools/:id - Fetch tool by ID
func GetToolByID(c *gin.Context) {
	var tool models.Tool
	id := c.Param("id")

	if err := database.DB.Preload("Tool").First(&tool, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, tool)
}

// POST /tools/ - Insert a new tool
func CreateTool(c *gin.Context) {
	// Parse form data (10 MB max)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	// Read fields from the form
	title := c.PostForm("title")
	description := c.PostForm("description")
	url := c.PostForm("url")

	// Handle file upload
	fileHeader, err := c.FormFile("cover")
	var coverPath string
	if err == nil {
		// Save the uploaded file to the "uploads" directory
		uploadDir := "./uploads"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		coverPath = filepath.Join(uploadDir, fileHeader.Filename)
		if err := c.SaveUploadedFile(fileHeader, coverPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save cover image"})
			return
		}
	}

	// Create new tool entry
	newTool := models.Tool{
		Title:       title,
		Description: description,
		URL:         url,
		Cover:       coverPath, // Save the relative path
	}

	if result := database.DB.Create(&newTool); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, newTool)
}

// PUT /tools/:id/edit - Edit a tool
func EditTool(c *gin.Context) {
	var newTool models.Tool
	if err := c.ShouldBindJSON(&newTool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var existingTool models.Tool
	if err := database.DB.Preload("Questionnaire").First(&existingTool, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Update only the fields sent in the request
	if err := database.DB.Model(&existingTool).Updates(newTool).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update answer"})
		return
	}
	c.JSON(http.StatusOK, existingTool)
}
