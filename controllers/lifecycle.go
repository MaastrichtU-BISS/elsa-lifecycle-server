package controllers

import (
	"net/http"
	"server/database"
	"server/models"
	"server/utils"

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
	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	if err := database.DB.Preload("Phases.Reflections").First(&lifecycles, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, lifecycles)
}
