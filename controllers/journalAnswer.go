package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /journalAnswers/:id - Fetch journalAnswers by ID
func GetJournalAnswerByID(c *gin.Context) {
	var answer models.JournalAnswer
	id := c.Param("id")

	if err := database.DB.Preload("Journal").First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// POST /journalAnswers - Insert a new journalAnswers
func CreateJournalAnswer(c *gin.Context) {
	var newAnswer models.JournalAnswer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newAnswer)
	c.JSON(http.StatusOK, newAnswer)
}

// PUT /journalAnswers/:id/edit - Edit an journalAnswer
func EditJournalAnswer(c *gin.Context) {
	var newAnswer models.JournalAnswer
	if err := c.ShouldBindJSON(&newAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var existingAnswer models.JournalAnswer
	if err := database.DB.Preload("Journal").First(&existingAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Update only the fields sent in the request
	if err := database.DB.Model(&existingAnswer).
		Updates(&newAnswer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	//fetch the updated answer
	var updatedAnswer models.JournalAnswer
	if err := database.DB.Preload("Journal").First(&updatedAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}
