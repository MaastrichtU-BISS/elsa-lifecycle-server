package controllers

import (
	"errors"
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// GET /furtherReflectionAnswers?rid=:rid - Fetch furtherReflectionAnswer by userId and reflectionId
func GetFurtherReflectionAnswerByUserIdAndReflectionID(c *gin.Context) {
	var answer models.FurtherReflectionAnswer
	rid := c.Query("rid")
	userId := c.GetString("user_id")

	result := database.DB.
		Preload("Reflection").
		Where("reflection_id = ? AND user_id = ?", rid, userId).
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

	newAnswer.UserID = uuid.MustParse(c.GetString("user_id"))
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
