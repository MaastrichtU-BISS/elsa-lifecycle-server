package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"server/database"
	"server/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
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

	if err := database.DB.Preload("Phases.Reflections").First(&lifecycles, id).Error; err != nil {
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
		Reflection            models.Reflection
		Answer                *models.ReflectionAnswer
		FurtherAnswer         *models.FurtherReflectionAnswer
		Recommendations       []models.Recommendation
		RecommendationAnswers map[uint]*models.RecommendationAnswer
	}

	type PhaseData struct {
		Phase              models.Phase
		ReflectionDataList []ReflectionData
	}

	// Optional reflectionId filter
	reflectionIDParam := c.Query("reflectionId")

	var phaseDataList []PhaseData

	for _, phase := range lifecycle.Phases {
		var reflectionDataList []ReflectionData

		for _, reflection := range phase.Reflections {
			if reflectionIDParam != "" && fmt.Sprintf("%d", reflection.ID) != reflectionIDParam {
				continue
			}
			// Get user's answer to this reflection
			answer := models.ReflectionAnswer{ReflectionID: reflection.ID, Reflection: reflection, UserID: userUUID, UpdatedAt: time.Now()}
			answerPtr := &answer

			// Get user's further reflection answer
			var furtherAnswer models.FurtherReflectionAnswer
			var furtherAnswerPtr *models.FurtherReflectionAnswer
			if err := database.DB.
				Where("reflection_id = ? AND user_id = ?", reflection.ID, userUUID).
				First(&furtherAnswer).Error; err == nil {
				furtherAnswerPtr = &furtherAnswer
			}

			// Get recommendations for this reflection
			var recommendations []models.Recommendation
			database.DB.
				Preload("Tool").
				Where("reflection_id = ?", reflection.ID).
				Find(&recommendations)

			// Get user's recommendation answers that are checked done
			recommendationAnswers := make(map[uint]*models.RecommendationAnswer)
			for _, rec := range recommendations {
				var recAnswer models.RecommendationAnswer
				if err := database.DB.
					Where("recommendation_id = ? AND user_id = ? AND checked_done = ?", rec.ID, userUUID, true).
					First(&recAnswer).Error; err == nil {
					recommendationAnswers[rec.ID] = &recAnswer
				}
			}

			reflectionDataList = append(reflectionDataList, ReflectionData{
				Reflection:            reflection,
				Answer:                answerPtr,
				FurtherAnswer:         furtherAnswerPtr,
				Recommendations:       recommendations,
				RecommendationAnswers: recommendationAnswers,
			})
		}

		if len(reflectionDataList) > 0 {
			phaseDataList = append(phaseDataList, PhaseData{
				Phase:              phase,
				ReflectionDataList: reflectionDataList,
			})
		}
	}

	// Create PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	if reflectionIDParam == "" {

		// Title
		pdf.SetFont("Arial", "B", 24)
		pdf.CellFormat(0, 15, stripMarkdown(lifecycle.Title), "", 1, "C", false, 0, "")
		pdf.Ln(5)

		// Username (without "User:" prefix)
		pdf.SetFont("Arial", "I", 11)
		pdf.CellFormat(0, 6, user.Email, "", 1, "L", false, 0, "")

		// Date and time of generation
		pdf.CellFormat(0, 6, fmt.Sprintf("Generated: %s", time.Now().Format("January 2, 2006 at 3:04 PM")), "", 1, "L", false, 0, "")
		pdf.Ln(8)

		// Description
		if lifecycle.Description != "" {
			pdf.SetFont("Arial", "", 12)
			renderMarkdownText(pdf, lifecycle.Description, 12)
			pdf.Ln(3)
		}

		// Introduction section
		if lifecycle.Introduction != "" {
			pdf.SetFont("Arial", "B", 16)
			pdf.CellFormat(0, 10, "Introduction", "", 1, "L", false, 0, "")
			renderMarkdownText(pdf, lifecycle.Introduction, 12)
			pdf.Ln(8)
		}
	}

	// Phases
	for _, phaseData := range phaseDataList {
		// Check page break before phase title
		if pdf.GetY() > 250 {
			pdf.AddPage()
		}

		// Phase Title
		pdf.SetFont("Arial", "B", 18)
		pdf.CellFormat(0, 10, cleanText(stripMarkdown(phaseData.Phase.Title)), "", 1, "L", false, 0, "")
		pdf.Ln(2)

		// Phase Description
		if phaseData.Phase.Description != "" {
			renderMarkdownText(pdf, phaseData.Phase.Description, 12)
			pdf.Ln(5)
		}

		// Reflections for this phase
		for _, reflData := range phaseData.ReflectionDataList {
			// Check page break before reflection title
			if pdf.GetY() > 250 {
				pdf.AddPage()
			}
			// Reflection Title
			pdf.SetFont("Arial", "B", 14)
			pdf.CellFormat(0, 8, cleanText(stripMarkdown(reflData.Reflection.Title)), "", 1, "L", false, 0, "")
			pdf.Ln(1)

			// Reflection Description
			if reflData.Reflection.Description != "" {
				renderMarkdownText(pdf, reflData.Reflection.Description, 11)
				pdf.Ln(2)
			}

			// Considerations
			if reflData.Reflection.Considerations != "" {
				// Check page break
				if pdf.GetY() > 250 {
					pdf.AddPage()
				}

				pdf.SetFont("Arial", "B", 11)
				pdf.CellFormat(0, 6, "In your answer, you might consider:", "", 1, "L", false, 0, "")

				// Parse considerations as JSON array
				var considerations []string
				if err := json.Unmarshal([]byte(reflData.Reflection.Considerations), &considerations); err == nil {
					pdf.SetFont("Arial", "", 11)
					for _, consideration := range considerations {
						pdf.MultiCell(0, 5, fmt.Sprintf("  - %s", cleanText(stripMarkdown(consideration))), "", "L", false)
					}
				} else {
					// Fallback if not JSON, just render as text
					renderMarkdownText(pdf, reflData.Reflection.Considerations, 11)
				}
				pdf.Ln(2)
			}

			// Answer
			if reflData.Answer != nil && reflData.Answer.Form != "" {
				// Check page break
				if pdf.GetY() > 250 {
					pdf.AddPage()
				}

				pdf.SetFont("Arial", "B", 12)
				pdf.CellFormat(0, 6, "Answer:", "", 1, "L", false, 0, "")
				// Parse and render reflection answer fields (without showing titles)
				renderFields(pdf, reflData.Answer.Form, 11, []FieldConfig{
					{Key: "free_text", Title: "", Style: ""},
					{Key: "get_recommendations", Title: "", Style: "I"},
				}, false)
				pdf.Ln(4)
			}

			// Recommended Tools Used (only those checked done by user)
			var usedRecommendations []models.Recommendation
			for _, rec := range reflData.Recommendations {
				if _, exists := reflData.RecommendationAnswers[rec.ID]; exists {
					usedRecommendations = append(usedRecommendations, rec)
				}
			}

			if len(usedRecommendations) > 0 {
				// Check page break
				if pdf.GetY() > 250 {
					pdf.AddPage()
				}

				pdf.SetFont("Arial", "B", 12)
				pdf.CellFormat(0, 6, "Recommended Tools Used:", "", 1, "L", false, 0, "")

				for _, rec := range usedRecommendations {
					pdf.SetFont("Arial", "", 10)
					// Title and description in black
					recText := fmt.Sprintf("  - %s", cleanText(stripMarkdown(rec.Tool.Title)))
					pdf.MultiCell(0, 5, recText, "", "L", false)

					// URL in light blue if present
					if rec.Tool.URL != "" {
						pdf.SetTextColor(70, 130, 180) // Light blue color (Steel Blue)
						pdf.SetFont("Arial", "U", 10)  // Underlined
						pdf.SetX(pdf.GetX() + 5)       // Indent 5mm from left margin
						pdf.MultiCell(0, 5, rec.Tool.URL, "", "L", false)
						pdf.SetTextColor(0, 0, 0)    // Reset to black
						pdf.SetFont("Arial", "", 10) // Reset font
					}
				}
				pdf.Ln(4)
			}

			// Further Reflections
			if reflData.FurtherAnswer != nil && reflData.FurtherAnswer.Form != "" {
				// Check page break
				if pdf.GetY() > 250 {
					pdf.AddPage()
				}

				pdf.SetFont("Arial", "B", 12)
				pdf.CellFormat(0, 6, "Further Reflections", "", 1, "L", false, 0, "")
				renderFields(pdf, reflData.FurtherAnswer.Form, 11, []FieldConfig{
					{Key: "further-reflection", Title: "", Style: ""},
				}, false)
				pdf.Ln(4)
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

// stripMarkdown removes basic markdown syntax for cleaner PDF output
func stripMarkdown(text string) string {
	// Remove markdown headers
	text = strings.ReplaceAll(text, "### ", "")
	text = strings.ReplaceAll(text, "## ", "")
	text = strings.ReplaceAll(text, "# ", "")

	// Remove bold/italic markers - be thorough
	text = strings.ReplaceAll(text, "**", "")
	text = strings.ReplaceAll(text, "__", "")
	// Only remove single * and _ if they're markers, not part of text
	text = strings.ReplaceAll(text, " * ", " ")
	text = strings.ReplaceAll(text, " _ ", " ")

	return strings.TrimSpace(text)
}

// cleanText removes problematic UTF-8 characters
func cleanText(text string) string {
	// Replace common UTF-8 encoding issues using hex codes
	replacements := map[string]string{
		"\xe2\x80\x99": "'",   // Smart apostrophe
		"\xe2\x80\x9c": "\"",  // Left double quote
		"\xe2\x80\x9d": "\"",  // Right double quote
		"\xe2\x80\x93": "-",   // En dash
		"\xe2\x80\x94": "-",   // Em dash
		"\xe2\x80\xa6": "...", // Ellipsis
		"\xc2\xa0":     " ",   // Non-breaking space
		"\xc2":         "",    // Stray Â character
	}

	for old, new := range replacements {
		text = strings.ReplaceAll(text, old, new)
	}

	return text
}

// FieldConfig defines a field to extract from form JSON
type FieldConfig struct {
	Key   string // JSON key to extract
	Title string // Optional title to display
	Style string // Font style: "", "I" for italic, "B" for bold
}

// renderFields parses JSON and extracts specified fields, optionally showing titles
func renderFields(pdf *gofpdf.Fpdf, formJSON string, fontSize float64, fields []FieldConfig, showTitles bool) {
	var formData map[string]interface{}
	if err := json.Unmarshal([]byte(formJSON), &formData); err != nil {
		// If not valid JSON, render as plain text
		pdf.SetFont("Arial", "", fontSize)
		pdf.MultiCell(0, 5, cleanText(stripMarkdown(formJSON)), "", "L", false)
		return
	}

	for _, field := range fields {
		// Handle both flat string values and nested @value structure
		var value string
		var ok bool

		// First try to get as string (for reflection answers)
		if value, ok = formData[field.Key].(string); !ok {
			// Try to get as object with @value.
			if fieldData, ok := formData[field.Key].(map[string]interface{}); ok {
				value, ok = fieldData["@value"].(string)
			}
		}

		if ok && value != "" {
			// Render title in bold if showTitles is true
			if showTitles && field.Title != "" {
				pdf.SetFont("Arial", "B", fontSize)
				pdf.MultiCell(0, 6, field.Title, "", "L", false)
				pdf.Ln(1)
			}

			// Render value with specified style
			style := field.Style
			if style == "" {
				style = ""
			}
			pdf.SetFont("Arial", style, fontSize)
			pdf.MultiCell(0, 5, cleanText(stripMarkdown(value)), "", "L", false)
			pdf.SetFont("Arial", "", fontSize) // Reset to normal
			pdf.Ln(4)
		}
	}
}

// renderMarkdownText renders text with basic markdown parsing
func renderMarkdownText(pdf *gofpdf.Fpdf, text string, fontSize float64) {
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			pdf.Ln(3)
			continue
		}

		// Handle headers
		if strings.HasPrefix(line, "### ") {
			pdf.SetFont("Arial", "B", fontSize+2)
			pdf.MultiCell(0, 5, strings.TrimPrefix(line, "### "), "", "L", false)
			pdf.SetFont("Arial", "", fontSize)
		} else if strings.HasPrefix(line, "## ") {
			pdf.SetFont("Arial", "B", fontSize+4)
			pdf.MultiCell(0, 6, strings.TrimPrefix(line, "## "), "", "L", false)
			pdf.SetFont("Arial", "", fontSize)
		} else if strings.HasPrefix(line, "# ") {
			pdf.SetFont("Arial", "B", fontSize+6)
			pdf.MultiCell(0, 7, strings.TrimPrefix(line, "# "), "", "L", false)
			pdf.SetFont("Arial", "", fontSize)
		} else if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
			// Bullet points - strip ** from them
			pdf.SetFont("Arial", "", fontSize)
			pdf.MultiCell(0, 5, cleanText("  "+stripMarkdown(line)), "", "L", false)
		} else {
			// Regular text - strip markdown and clean
			pdf.SetFont("Arial", "", fontSize)
			pdf.MultiCell(0, 5, cleanText(stripMarkdown(line)), "", "L", false)
		}
	}
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

type LastUpdatedItem struct {
	Type           string    `json:"type"`
	Title          string    `json:"title"`
	ToolID         uint      `json:"toolId,omitempty"`
	LifecycleID    uint      `json:"lifecycleId"`
	LifecycleTitle string    `json:"lifecycleTitle"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

const (
	ActivityReflection        = "reflection"
	ActivityFurtherReflection = "further_reflection"
	ActivityRecommendation    = "recommendation"
)

type LatestLifecycleResults struct {
	Reflection        *models.ReflectionAnswer
	FurtherReflection *models.FurtherReflectionAnswer
	Recommendation    *models.RecommendationAnswer
}

func FetchFirst[T any](query *gorm.DB, dest *T) error {
	return query.First(dest).Error
}

func PreloadReflectionLifecycle(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Reflection").
		Preload("Reflection.Phase").
		Preload("Reflection.Phase.Lifecycle")
}

func LatestByUser(db *gorm.DB, userID string) *gorm.DB {
	return db.
		Where("user_id = ?", userID).
		Order("updated_at DESC")
}

func PreloadRecommendationLifecycle(db *gorm.DB) *gorm.DB {
	return db.
		Preload("Recommendation").
		Preload("Recommendation.Reflection").
		Preload("Recommendation.Reflection.Phase").
		Preload("Recommendation.Reflection.Phase.Lifecycle")
}

func ToLastUpdatedItemReflection(a models.ReflectionAnswer) LastUpdatedItem {
	return LastUpdatedItem{
		Type:           ActivityReflection,
		Title:          a.Reflection.Title,
		LifecycleID:    a.Reflection.Phase.LifecycleID,
		LifecycleTitle: a.Reflection.Phase.Lifecycle.Title,
		UpdatedAt:      a.UpdatedAt,
	}
}

func ToLastUpdatedItemFurtherReflection(a models.FurtherReflectionAnswer) LastUpdatedItem {
	return LastUpdatedItem{
		Type:           ActivityFurtherReflection,
		Title:          a.Reflection.Title,
		LifecycleID:    a.Reflection.Phase.LifecycleID,
		LifecycleTitle: a.Reflection.Phase.Lifecycle.Title,
		UpdatedAt:      a.UpdatedAt,
	}
}

func ToLastUpdatedItemRecommendation(a models.RecommendationAnswer) LastUpdatedItem {
	return LastUpdatedItem{
		Type:           ActivityRecommendation,
		Title:          a.Recommendation.Reflection.Title,
		ToolID:         a.Recommendation.ToolID,
		LifecycleID:    a.Recommendation.Reflection.Phase.LifecycleID,
		LifecycleTitle: a.Recommendation.Reflection.Phase.Lifecycle.Title,
		UpdatedAt:      a.UpdatedAt,
	}
}

func PickLatest(current *LastUpdatedItem, candidate LastUpdatedItem) *LastUpdatedItem {
	if current == nil || candidate.UpdatedAt.After(current.UpdatedAt) {
		item := candidate
		return &item
	}
	return current
}

// GET /progress/last-updated - Get the last updated answer (reflection, further reflection, or recommendation) for the user
func GetLastUpdatedLifecycleItemForUser(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	if _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id must be a valid UUID"})
		return
	}

	results := LatestLifecycleResults{}
	db := database.DB.WithContext(c.Request.Context())

	// Use sequential queries in test mode to avoid SQLite in-memory locking issues
	if gin.Mode() == gin.TestMode {
		// ReflectionAnswer
		var a models.ReflectionAnswer
		err := FetchFirst(LatestByUser(PreloadReflectionLifecycle(db), userID), &a)
		if err == nil {
			results.Reflection = &a
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		// FurtherReflectionAnswer
		var b models.FurtherReflectionAnswer
		err = FetchFirst(LatestByUser(PreloadReflectionLifecycle(db), userID), &b)
		if err == nil {
			results.FurtherReflection = &b
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		// RecommendationAnswer
		var cRec models.RecommendationAnswer
		err = FetchFirst(LatestByUser(PreloadRecommendationLifecycle(db), userID), &cRec)
		if err == nil {
			results.Recommendation = &cRec
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}
	} else {
		// Default: parallel queries (production)
		g, ctx := errgroup.WithContext(c.Request.Context())
		db := database.DB.WithContext(ctx)

		g.Go(func() error {
			var a models.ReflectionAnswer
			err := FetchFirst(LatestByUser(PreloadReflectionLifecycle(db), userID), &a)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil
				}
				return err
			}
			results.Reflection = &a
			return nil
		})

		g.Go(func() error {
			var a models.FurtherReflectionAnswer
			err := FetchFirst(LatestByUser(PreloadReflectionLifecycle(db), userID), &a)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil
				}
				return err
			}
			results.FurtherReflection = &a
			return nil
		})

		g.Go(func() error {
			var a models.RecommendationAnswer
			err := FetchFirst(LatestByUser(PreloadRecommendationLifecycle(db), userID), &a)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil
				}
				return err
			}
			results.Recommendation = &a
			return nil
		})

		if err := g.Wait(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}
	}

	var latest *LastUpdatedItem
	if results.Reflection != nil {
		latest = PickLatest(latest, ToLastUpdatedItemReflection(*results.Reflection))
	}
	if results.FurtherReflection != nil {
		latest = PickLatest(latest, ToLastUpdatedItemFurtherReflection(*results.FurtherReflection))
	}
	if results.Recommendation != nil {
		latest = PickLatest(latest, ToLastUpdatedItemRecommendation(*results.Recommendation))
	}

	c.JSON(http.StatusOK, gin.H{"data": latest})
}
