package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// GET /reflectionAnswers?rid=:rid - Fetch reflectionAnswer by ID
func GetReflectionAnswerByUserIdAndReflectionID(c *gin.Context) {
	var answer models.ReflectionAnswer
	rid := c.Query("rid")
	userId := c.GetString("user_id") // Assuming user ID is stored in context after authentication

	if err := database.DB.Preload("Reflection").
		Where("reflection_id = ? AND user_id = ?", rid, userId).
		First(&answer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Items not found"})
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

	newAnswer.UserID = uuid.MustParse(c.GetString("user_id")) // Assuming user ID is stored in context after authentication
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
