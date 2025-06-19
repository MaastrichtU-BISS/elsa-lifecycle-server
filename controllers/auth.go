// controllers/auth.go
package controllers

import (
	"server/database"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if input.Password != input.ConfirmPassword {
		c.JSON(400, gin.H{"error": "Passwords do not match"})
		return
	}

	hashed, _ := utils.HashPassword(input.Password)
	user := models.User{Email: input.Email, PasswordHash: hashed}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	token, _ := utils.GenerateJWT(user.ID)
	c.JSON(200, gin.H{"token": token})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		c.JSON(401, gin.H{"error": "Wrong password"})
		return
	}

	token, _ := utils.GenerateJWT(user.ID)
	c.JSON(200, gin.H{"token": token})
}
