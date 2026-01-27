package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"server/database"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

// GET /lifecycles - Fetch all lifecycles
func GetLifecycles(c *gin.Context) {
	var lifecycles []models.Lifecycle
	database.DB.Find(&lifecycles)
	c.JSON(http.StatusOK, lifecycles)
}

// GET /lifecycles/:id - Fetch lifecycle by ID
func GetLifecyclesByID(c *gin.Context) {
	var lifecycles models.Lifecycle
	id := c.Param("id")

	if err := database.DB.Preload("Phases.Reflections").Preload("Phases.Journal").First(&lifecycles, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, lifecycles)
}

// GET /lifecycles/:id/pdf - Generate PDF for a lifecycle with user's answers
func GenerateLifecyclePDF(c *gin.Context) {
	lifecycleID := c.Param("id")
	userID := c.GetString("user_id") // pulled from JWT in middleware

	// Parse user UUID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch user info
	var user models.User
	if err := database.DB.Where("id = ?", userUUID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Fetch lifecycle with all related data
	var lifecycle models.Lifecycle
	if err := database.DB.
		Preload("Phases.Reflections").
		Where("id = ?", lifecycleID).
		First(&lifecycle).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lifecycle not found"})
		return
	}

	// For each reflection, get the user's answer and recommendations
	type ReflectionData struct {
		Reflection      models.Reflection
		Answer          *models.ReflectionAnswer
		Recommendations []models.Recommendation
	}

	phaseDataMap := make(map[uint][]ReflectionData)

	for _, phase := range lifecycle.Phases {
		var reflectionDataList []ReflectionData

		for _, reflection := range phase.Reflections {
			// Get user's answer to this reflection
			var answer models.ReflectionAnswer
			answerErr := database.DB.
				Where("reflection_id = ? AND user_id = ?", reflection.ID, userUUID).
				First(&answer).Error

			var answerPtr *models.ReflectionAnswer
			if answerErr == nil {
				answerPtr = &answer
			}

			// Get recommendations for this reflection
			var recommendations []models.Recommendation
			database.DB.
				Preload("Tool").
				Where("reflection_id = ?", reflection.ID).
				Find(&recommendations)

			reflectionDataList = append(reflectionDataList, ReflectionData{
				Reflection:      reflection,
				Answer:          answerPtr,
				Recommendations: recommendations,
			})
		}

		phaseDataMap[phase.ID] = reflectionDataList
	}

	// Create PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Title
	pdf.SetFont("Arial", "B", 24)
	pdf.CellFormat(0, 15, lifecycle.Title, "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Description
	if lifecycle.Description != "" {
		pdf.SetFont("Arial", "", 12)
		pdf.MultiCell(0, 6, lifecycle.Description, "", "L", false)
		pdf.Ln(3)
	}

	// Username
	pdf.SetFont("Arial", "I", 11)
	pdf.CellFormat(0, 6, fmt.Sprintf("User: %s", user.Email), "", 1, "L", false, 0, "")

	// Date and time of generation
	pdf.CellFormat(0, 6, fmt.Sprintf("Generated: %s", time.Now().Format("January 2, 2006 at 3:04 PM")), "", 1, "L", false, 0, "")
	pdf.Ln(8)

	// Welcome section
	if lifecycle.Welcome != "" {
		pdf.SetFont("Arial", "B", 16)
		pdf.CellFormat(0, 10, "Welcome", "", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		pdf.MultiCell(0, 6, lifecycle.Welcome, "", "L", false)
		pdf.Ln(5)
	}

	// Introduction section
	if lifecycle.Introduction != "" {
		pdf.SetFont("Arial", "B", 16)
		pdf.CellFormat(0, 10, "Introduction", "", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		pdf.MultiCell(0, 6, lifecycle.Introduction, "", "L", false)
		pdf.Ln(8)
	}

	// Phases
	for _, phase := range lifecycle.Phases {
		// Phase Title
		pdf.SetFont("Arial", "B", 18)
		pdf.CellFormat(0, 10, phase.Title, "", 1, "L", false, 0, "")
		pdf.Ln(2)

		// Phase Description
		if phase.Description != "" {
			pdf.SetFont("Arial", "", 12)
			pdf.MultiCell(0, 6, phase.Description, "", "L", false)
			pdf.Ln(5)
		}

		// Reflections for this phase
		reflectionDataList := phaseDataMap[phase.ID]
		for _, reflData := range reflectionDataList {
			// Reflection Title
			pdf.SetFont("Arial", "B", 14)
			pdf.CellFormat(0, 8, reflData.Reflection.Title, "", 1, "L", false, 0, "")
			pdf.Ln(1)

			// Reflection Description
			if reflData.Reflection.Description != "" {
				pdf.SetFont("Arial", "", 11)
				pdf.MultiCell(0, 5, reflData.Reflection.Description, "", "L", false)
				pdf.Ln(2)
			}

			// Consideration
			if reflData.Reflection.Considerations != "" {
				pdf.SetFont("Arial", "I", 11)
				pdf.MultiCell(0, 5, fmt.Sprintf("Considerations: %s", reflData.Reflection.Considerations), "", "L", false)
				pdf.Ln(2)
			}

			// Answer
			if reflData.Answer != nil && reflData.Answer.Form != "" {
				pdf.SetFont("Arial", "", 11)
				pdf.MultiCell(0, 5, fmt.Sprintf("Answer: %s", reflData.Answer.Form), "", "L", false)
				pdf.Ln(3)
			}

			// Recommendations
			if len(reflData.Recommendations) > 0 {
				pdf.SetFont("Arial", "B", 12)
				pdf.CellFormat(0, 6, "Recommendations:", "", 1, "L", false, 0, "")

				for _, rec := range reflData.Recommendations {
					pdf.SetFont("Arial", "", 10)
					recText := fmt.Sprintf("- %s", rec.Tool.Title)
					if rec.Tool.Description != "" {
						recText += fmt.Sprintf(": %s", rec.Tool.Description)
					}
					if rec.Tool.URL != "" {
						recText += fmt.Sprintf(" (%s)", rec.Tool.URL)
					}
					pdf.MultiCell(0, 5, recText, "", "L", false)
				}
				pdf.Ln(2)
			}

			pdf.Ln(4)
		}

		pdf.Ln(3)
	}

	// Generate PDF to buffer
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF", "details": err.Error()})
		return
	}

	// Set headers for file download
	filename := fmt.Sprintf("%s.pdf", lifecycle.Title)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	c.Header("Content-Type", "application/pdf")
	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}

// GET lifecycles/:id/phases - Fetch all phases given a lifecycle ID
// func GetPhases(c *gin.Context) {
// 	var phases models.Phase
// 	id := c.Param("id")

// 	if err := database.DB.Preload("Reflections").
// 		Where("lifecycle_id = ?", id).
// 		Find(&phases).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, phases)
// }
