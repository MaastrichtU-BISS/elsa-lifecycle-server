// controllers/auth.go
package controllers

import (
	"server/database"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	if err := utils.ValidatePassword(input.Password); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not hash password"})
		return
	}

	// Generate the session token before inserting the user, so a signing failure
	// cannot leave behind an account the client never received a token for.
	user := models.User{ID: uuid.New(), Email: input.Email, PasswordHash: hashed}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create session"})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

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

	if err := utils.ValidatePassword(input.Password); err != nil {
		c.JSON(401, gin.H{"error": "Wrong password"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		c.JSON(401, gin.H{"error": "Wrong password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create session"})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
