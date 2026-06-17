package controllers

import (
	"errors"
	"net/http"
	"server/database"
	"server/models"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET /reflectionAnswers/:id - Fetch reflectionAnswer by ID
func GetReflectionAnswerByID(c *gin.Context) {
	var answer models.ReflectionAnswer
	id := c.Param("id")

	if err := database.DB.Preload("Reflection").First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// GET /reflectionAnswers/:id?jid=:jid
func GetReflectionAnswerByJournalIdAndReflectionID(c *gin.Context) {
	var answer models.ReflectionAnswer
	jid := c.Query("jid")
	rid := c.Query("id")
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
	if err := utils.CheckJournalAuthentication(strconv.FormatUint(uint64(newAnswer.JournalID), 10), userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newAnswer)
	c.JSON(http.StatusOK, newAnswer)
}

// PUT /reflectionAnswers/:id/edit - Edit an reflectionAnswer
func EditReflectionAnswer(c *gin.Context) {
	var newAnswer models.ReflectionAnswer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("user_id")
	// Validate journal ownership and existence
	if err := utils.CheckJournalAuthentication(strconv.FormatUint(uint64(newAnswer.JournalID), 10), userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var existingAnswer models.ReflectionAnswer
	if err := database.DB.Preload("Reflection").First(&existingAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Update only the fields sent in the request
	if err := database.DB.Model(&existingAnswer).
		Select("Form", "BinaryEvaluation").
		Updates(&newAnswer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update answer"})
		return
	}

	//fetch the updated answer
	var updatedAnswer models.ReflectionAnswer
	if err := database.DB.Preload("Reflection").First(&updatedAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}
