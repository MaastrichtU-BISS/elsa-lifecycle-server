package controllers

import (
	"net/http"
	"os"

	"server/seeder"

	"github.com/gin-gonic/gin"
)

func ResetDatabase(c *gin.Context) {
	if os.Getenv("APP_ENV") != "test" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "not allowed outside test environment",
		})
		return
	}

	expectedSecret := os.Getenv("E2E_RESET_SECRET")
	givenSecret := c.GetHeader("x-e2e-reset-secret")

	if expectedSecret == "" || givenSecret != expectedSecret {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	if err := seeder.RequireTestEnvironment(); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := seeder.ResetAndSeedDatabase(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
