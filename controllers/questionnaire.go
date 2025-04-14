package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /answers/questionnaires_id - Fetch all questionnaires
func GetQuestionnaires(c *gin.Context) {
	var questionnaires []models.Questionnaire
	database.DB.Find(&questionnaires)
	c.JSON(http.StatusOK, questionnaires)
}

// GET /answers/id - Fetch answer by ID
func GetQuestionnaireByID(c *gin.Context) {
	var questionnaire models.Questionnaire
	id := c.Param("id")

	if err := database.DB.Preload("Answers").First(&questionnaire, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, questionnaire)
}

// POST /answers - Insert a new answer
func CreateAnswer(c *gin.Context) {
	var newAnswer models.Answer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newAnswer)
	c.JSON(http.StatusOK, newAnswer)
}
