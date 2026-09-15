package controllers

import (
	"net/http"
	"server/database"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// GET /phases/:id - Fetch phase by id
func GetPhaseById(c *gin.Context) {
	var phase models.Phase
	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	if err := database.DB.Preload("Reflections", orderBy("reflections.id")).
		First(&phase, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, phase)
}

// GET /phases/:id/reflections - Fetch all reflections that belong to a phase
// func getReflectionsByPhaseID(c *gin.Context) ([]models.Reflection, error) {
// 	var reflections []models.Reflection
// 	phaseID := c.Param("id")
// 	if err := database.DB.Where("phase_id = ?", phaseID).Find(&reflections).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
// 		return nil, err
// 	}
// 	return reflections, nil
// }
