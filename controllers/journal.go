package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// GET /journals/:id - Fetch journal by ID
func GetJournalByID(c *gin.Context) {
	var journal models.Journal
	id := c.Param("id")

	if err := database.DB.First(&journal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, journal)
}

// DELETE /journals/:id - Delete journal by ID
func DeleteJournal(c *gin.Context) {
	var journal models.Journal
	id := c.Param("id")

	if err := database.DB.First(&journal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := database.DB.Delete(&journal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Item deleted successfully"})
}

// POST /journals - Insert a new journal
func CreateJournal(c *gin.Context) {
	var newJournal models.Journal
	if err := c.ShouldBindJSON(&newJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&newJournal)
	c.JSON(http.StatusOK, newJournal)
}

// GET /journals/:id/answers - Fetch all answers that belong to a journal
func GetAnswers(c *gin.Context) {
	var answers []models.JournalAnswer
	journal_id := c.Param("id")

	if err := database.DB.Where("journal_id = ?", journal_id).
		Find(&answers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answers)
}
