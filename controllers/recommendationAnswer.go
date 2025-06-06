package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"server/database"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /recommendationAnswers/:id - Fetch recommendationAnswers by ID
func GetRecommendationAnswerByID(c *gin.Context) {
	var answer models.RecommendationAnswer
	id := c.Param("id")

	if err := database.DB.Preload("Recommendation").First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// POST /recommendationAnswers - Insert a new recommendationAnswers
func CreateRecommendationAnswer(c *gin.Context) {
	// Parse form data (10 MB max)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	// Read fields from the form
	form := c.PostForm("form")
	recommendationId, err := strconv.ParseUint(c.PostForm("recommendationId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse recommendationId to uint64"})
	}

	// Handle file upload
	fileHeader, err := c.FormFile("file")
	var filePath string
	if err == nil {
		// Save the uploaded file to the "uploads" directory
		uploadDir := "./uploads/recommendation_answers"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload/recommendations directory"})
			return
		}

		filePath = filepath.Join(uploadDir, fileHeader.Filename)
		if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
	}

	// Create new recommendation_answer entry
	newRecommendationAnswer := models.RecommendationAnswer{
		Form:             form,
		RecommendationID: uint(recommendationId),
		File:             filePath, // Save the relative path
	}

	if result := database.DB.Create(&newRecommendationAnswer); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, newRecommendationAnswer)
}

// PUT /recommendationAnswers/:id/edit - Edit an recommendationAnswer
func EditRecommendationAnswer(c *gin.Context) {
	var newAnswer models.RecommendationAnswer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var existingAnswer models.RecommendationAnswer
	if err := database.DB.Preload("Recommendation").First(&existingAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Update only the fields sent in the request
	if err := database.DB.Model(&existingAnswer).
		Updates(&newAnswer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	//fetch the updated answer
	var updatedAnswer models.RecommendationAnswer
	if err := database.DB.Preload("Recommendation").First(&updatedAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}
