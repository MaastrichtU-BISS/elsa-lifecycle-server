package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /questionnaires - Fetch all questionnaires
func GetQuestionnaires(c *gin.Context) {
	var questionnaires []models.Questionnaire
	database.DB.Find(&questionnaires)
	c.JSON(http.StatusOK, questionnaires)
}

// GET /questionnaires/id - Fetch questionnaire by ID
func GetQuestionnaireByID(c *gin.Context) {
	var questionnaire models.Questionnaire
	id := c.Param("id")

	if err := database.DB.First(&questionnaire, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, questionnaire)
}

// POST /questionnaires - Insert a new questionnaire
func CreateQuestionnaire(c *gin.Context) {
	var newQuestionnaire models.Questionnaire
	if err := c.ShouldBindJSON(&newQuestionnaire); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newQuestionnaire)
	c.JSON(http.StatusOK, newQuestionnaire)
}
