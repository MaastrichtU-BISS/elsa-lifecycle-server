package controllers

import (
	"errors"
	"net/http"
	"server/database"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET /reflectionAnswers/:id - Fetch reflectionAnswer by ID
func GetReflectionAnswerByID(c *gin.Context) {
	var answer models.ReflectionAnswer
	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	if err := database.DB.Preload("Reflection").First(&answer, id).Error; err != nil {
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

// GET /reflectionAnswers?rid=:rid&jid=:jid
func GetReflectionAnswerByJournalIdAndReflectionID(c *gin.Context) {
	var answer models.ReflectionAnswer
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
	userId := c.GetString("user_id")

	// Validate journal ownership and existence
	if err := utils.CheckJournalAuthentication(jid, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.
		Preload("Reflection").
		Where("reflection_id = ? AND journal_id = ?", rid, jid).
		First(&answer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, nil)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// POST /reflectionAnswers - Insert a new reflectionAnswer
func CreateReflectionAnswer(c *gin.Context) {

	var newAnswer models.ReflectionAnswer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("user_id")

	// Validate journal ownership and existence
	if err := utils.CheckJournalAuthentication(newAnswer.JournalID, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&newAnswer).Error; err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Reflection or journal does not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}
	touchJournal(newAnswer.JournalID)
	c.JSON(http.StatusOK, newAnswer)
}

// PUT /reflectionAnswers/:id/edit - Edit an reflectionAnswer
func EditReflectionAnswer(c *gin.Context) {
	var newAnswer models.ReflectionAnswer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}
	var existingAnswer models.ReflectionAnswer
	if err := database.DB.Preload("Reflection").First(&existingAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	userId := c.GetString("user_id") // Assuming user ID is stored in context after authentication

	// Validate journal ownership and existence before updating
	if err := utils.CheckJournalAuthentication(existingAnswer.JournalID, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// Update only the fields sent in the request
	if err := database.DB.Model(&existingAnswer).
		Select("Form", "BinaryEvaluation").
		Updates(&newAnswer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update answer"})
		return
	}

	touchJournal(existingAnswer.JournalID)

	//fetch the updated answer
	var updatedAnswer models.ReflectionAnswer
	if err := database.DB.Preload("Reflection").First(&updatedAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}
