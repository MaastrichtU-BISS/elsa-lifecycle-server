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
	getRecommendations := c.Query("getRecommendations") == "true"

	if !getRecommendations {
		c.JSON(http.StatusBadRequest, gin.H{"error": "getRecommendations query parameter must be 'true'"})
		return
	}

	if err := database.DB.Preload("Tool").
		Where("reflection_id = ?", reflectionId).
		Find(&recommendations).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Items not found"})
		return
	}

	c.JSON(http.StatusOK, recommendations)
}
