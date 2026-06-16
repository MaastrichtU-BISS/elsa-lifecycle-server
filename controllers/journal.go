package controllers

import (
	"errors"
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GET /journals/:id - Fetch journal by ID
func GetJournalByID(c *gin.Context) {
	var journal models.Journal
	id := c.Param("id")

	if err := database.DB.Preload("User").Preload("Lifecycle").First(&journal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, journal)
}

// GET /journals?userId=:userId&lifecycleId=:lifecycleId - Fetch journal by user ID and lifecycle ID
func GetJournalByUserIdAndLifecycleID(c *gin.Context) {
	var journal models.Journal
	userId := c.Query("userId")
	lifecycleId := c.Query("lifecycleId")

	result := database.DB.
		Preload("User").
		Preload("Lifecycle").
		Where("user_id = ? AND lifecycle_id = ?", userId, lifecycleId).
		First(&journal)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Not found — return null or empty response, not 404
			c.JSON(http.StatusOK, nil)
			return
		}
		// Some other DB error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item"})
		return
	}

	c.JSON(http.StatusOK, journal)
}

// POST /journals - Insert a new journal
func CreateJournal(c *gin.Context) {
	var newJournal models.Journal
	if err := c.ShouldBindJSON(&newJournal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newJournal.UserID = uuid.MustParse(c.GetString("user_id")) // Assuming user ID is stored in context after authentication
	database.DB.Create(&newJournal)
	c.JSON(http.StatusOK, newJournal)
}
