package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /lifecycles - Fetch all lifecycles
func GetLifecycles(c *gin.Context) {
	var lifecycles []models.Lifecycle
	database.DB.Find(&lifecycles)
	c.JSON(http.StatusOK, lifecycles)
}

// GET /lifecycles/:id - Fetch lifecycle by ID
func GetLifecyclesByID(c *gin.Context) {
	var lifecycles models.Lifecycle
	id := c.Param("id")

	if err := database.DB.Preload("Phases.Reflection.Answers").Preload("Phases.Journal.Answers").First(&lifecycles, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, lifecycles)
}

// GET lifecycles/:id/phases - Fetch all phases given a lifecycle ID
func GetPhases(c *gin.Context) {
	var phases models.Phase
	id := c.Param("id")

	if err := database.DB.Preload("Reflection").
		Where("lifecycle_id = ?", id).
		Find(&phases).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, phases)
}
