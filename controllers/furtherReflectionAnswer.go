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

// GET /furtherReflectionAnswers/:id - Fetch furtherReflectionAnswer by ID
func GetFurtherReflectionAnswerByID(c *gin.Context) {
	var answer models.FurtherReflectionAnswer
	id := c.Param("id")

	if err := database.DB.Preload("Reflection").First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// GET /furtherReflectionAnswers/:id?jid=:jid - Fetch furtherReflectionAnswer by journalId and reflectionId
func GetFurtherReflectionAnswerByJournalIdAndReflectionID(c *gin.Context) {
	var answer models.FurtherReflectionAnswer
	frid := c.Query("id")
	jid := c.Query("jid")
	userId := c.GetString("user_id")

	// Validate journal ownership and existence
	err := utils.CheckJournalAuthentication(jid, userId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.
		Preload("Reflection").
		Where("reflection_id = ? AND journal_id = ?", frid, jid).
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

// POST /furtherReflectionAnswers - Insert a new furtherReflectionAnswer
func CreateFurtherReflectionAnswer(c *gin.Context) {
	var newAnswer models.FurtherReflectionAnswer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("user_id")
	// Validate journal ownership and existence
	err := utils.CheckJournalAuthentication(strconv.FormatUint(uint64(newAnswer.JournalID), 10), userId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newAnswer)
	c.JSON(http.StatusOK, newAnswer)
}

// PUT /furtherReflectionAnswers/:id/edit - Edit a furtherReflectionAnswer
func EditFurtherReflectionAnswer(c *gin.Context) {
	var newAnswer models.FurtherReflectionAnswer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("user_id")
	// Validate journal ownership and existence
	err := utils.CheckJournalAuthentication(strconv.FormatUint(uint64(newAnswer.JournalID), 10), userId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var existingAnswer models.FurtherReflectionAnswer
	if err := database.DB.Preload("Reflection").First(&existingAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := database.DB.Model(&existingAnswer).
		Select("Form").
		Updates(&newAnswer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update answer"})
		return
	}

	var updatedAnswer models.FurtherReflectionAnswer
	if err := database.DB.Preload("Reflection").First(&updatedAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}
