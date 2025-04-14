package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /questionnaires/:id/answers - Fetch all answers that belong to a questionnaire
func GetAnswers(c *gin.Context) {
	var answers []models.Answer
	questionnaire_id := c.Param("id")

	database.DB.Where("questionnaire_id = ?", questionnaire_id).Find(&answers)
	c.JSON(http.StatusOK, answers)
}

// GET /amswers/id - Fetch questionnaire by ID
func GetAnswerByID(c *gin.Context) {
	var answer models.Answer
	id := c.Param("id")

	if err := database.DB.Preload("Questionnaire").First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// POST /questionnaires/:id/answers - Insert a new questionnaire
func CreateQuestionnaire(c *gin.Context) {
	var newQuestionnaire models.Questionnaire
	if err := c.ShouldBindJSON(&newQuestionnaire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newQuestionnaire)
	c.JSON(http.StatusOK, newQuestionnaire)
}
