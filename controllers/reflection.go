package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /reflections/:id - Fetch reflection by ID
func GetReflectionByID(c *gin.Context) {
	var reflection models.Reflection
	id := c.Param("id")

	if err := database.DB.Preload("ReflectionAnswers").First(&reflection, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, reflection)
}

// DELETE /reflections/:id - Delete reflection by ID
func DeleteReflection(c *gin.Context) {
	var reflection models.Reflection
	id := c.Param("id")

	if err := database.DB.First(&reflection, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := database.DB.Delete(&reflection).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reflection"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Reflection deleted successfully"})
}

// POST /reflections - Insert a new reflection
func CreateReflection(c *gin.Context) {
	var newReflection models.Reflection
	if err := c.ShouldBindJSON(&newReflection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newReflection)
	c.JSON(http.StatusOK, newReflection)
}

// GET /reflections/:id/answers - Fetch all answers that belong to a reflection
func GetReflectionAnswers(c *gin.Context) {
	var answers []models.ReflectionAnswer
	reflection_id := c.Param("id")

	if err := database.DB.
		Where("reflection_id = ?", reflection_id).
		Find(&answers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answers)
}
