package controllers

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"server/database"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET /recommendationAnswers/:id - Fetch recommendationAnswers by ID
func GetRecommendationAnswerByID(c *gin.Context) {
	var answer models.RecommendationAnswer
	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	if err := database.DB.Preload("Recommendation.Tool").First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	userId := c.GetString("user_id") // Assuming user ID is stored in context after authentication

	// Validate journal ownership and existence
	if err := utils.CheckJournalAuthentication(answer.JournalID, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// GET /recommendationAnswers?rid=:rid&jid=:jid - Fetch recommendationAnswer by journalId and recommendationId
func GetRecommendationAnswerByJournalIdAndRecommendationID(c *gin.Context) {
	var answer models.RecommendationAnswer
	rid, err := utils.ParseID(c.Query("rid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rid must be a positive integer"})
		return
	}
	jid, err := utils.ParseID(c.Query("jid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jid must be a positive integer"})
		return
	}
	userId := c.GetString("user_id") // Assuming user ID is stored in context after authentication

	// Validate journal ownership and existence
	if err := utils.CheckJournalAuthentication(jid, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.
		Preload("Recommendation.Tool").
		Where("recommendation_id = ? AND journal_id = ?", rid, jid).
		First(&answer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Not found — return null or empty response, not 404
			c.JSON(http.StatusOK, nil)
			return
		}
		// Some other DB error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item"})
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
	recommendationId, err := utils.ParseID(c.PostForm("recommendationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recommendationId must be a positive integer"})
		return
	}

	journalId, err := utils.ParseID(c.PostForm("journalId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "journalId must be a positive integer"})
		return
	}

	userId := c.GetString("user_id")
	// Validate journal ownership and existence before saving anything
	if err := utils.CheckJournalAuthentication(journalId, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// Parse checked_done boolean field
	checkedDone := c.PostForm("checked_done") == "true"

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
		RecommendationID: recommendationId,
		JournalID:        journalId,
		File:             filePath, // Save the relative path
		CheckedDone:      checkedDone,
	}

	// Step 1: Create the record
	if err := database.DB.Create(&newRecommendationAnswer).Error; err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Recommendation or journal does not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Step 2: Reload with preloads
	if err := database.DB.Preload("Recommendation.Tool").
		First(&newRecommendationAnswer, newRecommendationAnswer.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newRecommendationAnswer)
}

// PUT /recommendationAnswers/:id/edit - Edit an recommendationAnswer
func EditRecommendationAnswer(c *gin.Context) {
	// Parse form data (10 MB max)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}
	var existingAnswer models.RecommendationAnswer
	if err := database.DB.First(&existingAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	userId := c.GetString("user_id")
	// Validate journal ownership and existence before saving anything
	if err := utils.CheckJournalAuthentication(existingAnswer.JournalID, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// Read fields from the form
	form := c.PostForm("form")

	// Parse checked_done boolean field
	checkedDone := c.PostForm("checked_done") == "true"

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
		Form:        form,
		File:        filePath, // Save the relative path
		CheckedDone: checkedDone,
	}

	// Update only the fields sent in the request
	if err := database.DB.Model(&existingAnswer).
		// Without Select, checkedDone would not be updated if false
		Select("form", "file", "checked_done").
		Updates(&newRecommendationAnswer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	//fetch the updated answer
	var updatedAnswer models.RecommendationAnswer
	if err := database.DB.Preload("Recommendation.Tool").First(&updatedAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}
