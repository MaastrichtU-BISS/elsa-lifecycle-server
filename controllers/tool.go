package controllers

import (
	"net/http"
	"server/database"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// GET /tools/:id/tools - Fetch all tools
func GetTools(c *gin.Context) {
	var tools []models.Tool
	database.DB.Order("id").Find(&tools)
	c.JSON(http.StatusOK, tools)
}

// GET /tools/:id - Fetch tool by ID
func GetToolByID(c *gin.Context) {
	var tool models.Tool
	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	if err := database.DB.First(&tool, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, tool)
}
