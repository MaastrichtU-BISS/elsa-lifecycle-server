package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /phases/:id - Fetch phase by id
func GetPhaseById(c *gin.Context) {
	var phase models.Phase
	id := c.Param("id")

	if err := database.DB.Preload("Reflections").Preload("Journal").
		First(&phase, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, phase)
}