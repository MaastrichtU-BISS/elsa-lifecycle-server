package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /recommendations/:questionnaireId?binaryEvaluation - Fetch recommendations by questionnaire ID and binaryEvaluation
func GetRecommendations(c *gin.Context) {
	var recommendations []models.Recommendation
	questionnaireId := c.Param("questionnaireId")
	binaryEvaluation := c.Query("binaryEvaluation")

	if err := database.DB.Preload("Tool").
		Where("questionnaire_id = ? AND binary_evaluation = ?", questionnaireId, binaryEvaluation).
		Find(&recommendations).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Items not found"})
		return
	}

	c.JSON(http.StatusOK, recommendations)
}
