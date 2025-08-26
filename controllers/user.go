// controllers/user.go
package controllers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID := c.GetString("user_id") // pulled from JWT in middleware

	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
		"id":    user.ID,
	})
}
