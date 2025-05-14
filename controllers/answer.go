package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /answers/:id/answers - Fetch all answers that belong to a questionnaire
func GetAnswers(c *gin.Context) {
	var answers []models.Answer
	questionnaire_id := c.Param("id")

	database.DB.Where("questionnaire_id = ?", questionnaire_id).Find(&answers)
	c.JSON(http.StatusOK, answers)
}

// GET /amswers/:id - Fetch answer by ID
func GetAnswerByID(c *gin.Context) {
	var answer models.Answer
	id := c.Param("id")

	if err := database.DB.Preload("Questionnaire").First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// POST /answers/:id/answers - Insert a new answer
func CreateAnswer(c *gin.Context) {
	var newAnswer models.Answer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newAnswer)
	c.JSON(http.StatusOK, newAnswer)
}

// PUT /answers/:id/edit - Edit an answer
func EditAnswer(c *gin.Context) {
	var newAnswer models.Answer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var existingAnswer models.Answer
	if err := database.DB.Preload("Questionnaire").First(&existingAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
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
	var updatedAnswer models.Answer
	if err := database.DB.Preload("Questionnaire").First(&updatedAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}
