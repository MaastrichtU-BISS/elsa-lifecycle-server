package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /recommendations/:reflectionId?binaryEvaluation - Fetch recommendations by reflection ID and binaryEvaluation
func GetRecommendations(c *gin.Context) {
	var recommendations []models.Recommendation
	reflectionId := c.Param("reflectionId")
	binaryEvaluation := c.Query("binaryEvaluation")

	if err := database.DB.Preload("Tool").
		Preload("Answers").
		Where("reflection_id = ? AND binary_evaluation = ?", reflectionId, binaryEvaluation).
		Find(&recommendations).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Items not found"})
		return
	}

	c.JSON(http.StatusOK, recommendations)
}
