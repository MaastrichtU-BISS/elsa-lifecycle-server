package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /lifecycles - Fetch all lifecycles
func GetAllLifecycles(c *gin.Context) {
	var lifecycles []models.Lifecycle
	database.DB.Find(&lifecycles)
	c.JSON(http.StatusOK, lifecycles)
}

// GET /lifecycles/:id - Fetch lifecycle by ID
func GetLifecycleByID(c *gin.Context) {
	var lifecycles models.Lifecycle
	id := c.Param("id")

	if err := database.DB.Preload("Phases.Reflections").First(&lifecycles, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, lifecycles)
}
