package controllers

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"server/database"
	"server/models"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GET /recommendationAnswers/:id - Fetch recommendationAnswers by ID
func GetRecommendationAnswerByID(c *gin.Context) {
	var answer models.RecommendationAnswer
	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	if err := database.DB.Preload("Recommendation.Tool").First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	userId := c.GetString("user_id") // Assuming user ID is stored in context after authentication

	// Validate journal ownership and existence
	if err := utils.CheckJournalAuthentication(answer.JournalID, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answer)
}

// GET /recommendationAnswers?rid=:rid&jid=:jid - Fetch recommendationAnswer by journalId and recommendationId
func GetRecommendationAnswerByJournalIdAndRecommendationID(c *gin.Context) {
	var answer models.RecommendationAnswer
	rid, err := utils.ParseID(c.Query("rid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rid must be a positive integer"})
		return
	}
	jid, err := utils.ParseID(c.Query("jid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jid must be a positive integer"})
		return
	}
	userId := c.GetString("user_id") // Assuming user ID is stored in context after authentication

	// Validate journal ownership and existence
	if err := utils.CheckJournalAuthentication(jid, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.
		Preload("Recommendation.Tool").
		Where("recommendation_id = ? AND journal_id = ?", rid, jid).
		First(&answer)

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

	c.JSON(http.StatusOK, answer)
}

// uploaded files are private: only served through DownloadRecommendationAnswerFile
const recommendationAnswerUploadDir = "uploads/recommendation_answers"

// saveRecommendationAnswerFile stores an upload as <uploadDir>/<random id>/<original name>,
// so files with the same name never overwrite each other and the name stays readable
func saveRecommendationAnswerFile(c *gin.Context) (path string, uploaded bool, err error) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return "", false, nil
	}

	name := filepath.Base(strings.ReplaceAll(fileHeader.Filename, "\\", "/"))
	if name == "." || name == "/" || name == ".." {
		name = "file"
	}

	dir := filepath.Join(recommendationAnswerUploadDir, uuid.NewString())
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return "", true, err
	}

	path = filepath.Join(dir, name)
	if err := c.SaveUploadedFile(fileHeader, path); err != nil {
		os.RemoveAll(dir)
		return "", true, err
	}
	return path, true, nil
}

// removeRecommendationAnswerFile deletes a stored upload, ignoring paths outside the upload directory
func removeRecommendationAnswerFile(path string) {
	if !isInRecommendationAnswerUploadDir(path) {
		return
	}
	dir := filepath.Dir(filepath.Clean(path))
	if dir == filepath.Clean(recommendationAnswerUploadDir) {
		// file saved before per-upload directories
		os.Remove(path)
		return
	}
	os.RemoveAll(dir)
}

func isInRecommendationAnswerUploadDir(path string) bool {
	rel, err := filepath.Rel(recommendationAnswerUploadDir, filepath.Clean(path))
	return err == nil && rel != "." && !strings.HasPrefix(rel, "..")
}

// POST /recommendationAnswers - Insert a new recommendationAnswers
func CreateRecommendationAnswer(c *gin.Context) {
	// Parse form data (10 MB max)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	// Read fields from the form
	form := c.PostForm("form")
	recommendationId, err := utils.ParseID(c.PostForm("recommendationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recommendationId must be a positive integer"})
		return
	}

	journalId, err := utils.ParseID(c.PostForm("journalId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "journalId must be a positive integer"})
		return
	}

	userId := c.GetString("user_id")
	// Validate journal ownership and existence before saving anything
	if err := utils.CheckJournalAuthentication(journalId, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// Parse checked_done boolean field
	checkedDone := c.PostForm("checked_done") == "true"

	// Handle file upload
	filePath, _, err := saveRecommendationAnswerFile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Create new recommendation_answer entry
	newRecommendationAnswer := models.RecommendationAnswer{
		Form:             form,
		RecommendationID: recommendationId,
		JournalID:        journalId,
		File:             filePath, // Save the relative path
		CheckedDone:      checkedDone,
	}

	// Step 1: Create the record
	if err := database.DB.Create(&newRecommendationAnswer).Error; err != nil {
		removeRecommendationAnswerFile(filePath)
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Recommendation or journal does not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	// Step 2: Reload with preloads
	if err := database.DB.Preload("Recommendation.Tool").
		First(&newRecommendationAnswer, newRecommendationAnswer.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item"})
		return
	}

	c.JSON(http.StatusOK, newRecommendationAnswer)
}

// PUT /recommendationAnswers/:id/edit - Edit an recommendationAnswer.
// Only the fields present in the request (form, file, checked_done) are updated.
func EditRecommendationAnswer(c *gin.Context) {
	// Parse form data (10 MB max)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}
	var existingAnswer models.RecommendationAnswer
	if err := database.DB.First(&existingAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	userId := c.GetString("user_id")
	// Validate journal ownership and existence before saving anything
	if err := utils.CheckJournalAuthentication(existingAnswer.JournalID, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// a map (not a struct) so false and empty values are written too
	updates := map[string]interface{}{}
	if form, ok := c.GetPostForm("form"); ok {
		updates["form"] = form
	}
	if checkedDone, ok := c.GetPostForm("checked_done"); ok {
		updates["checked_done"] = checkedDone == "true"
	}

	filePath, uploaded, err := saveRecommendationAnswerFile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	if uploaded {
		updates["file"] = filePath
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nothing to update: send form, file or checked_done"})
		return
	}

	// Updates writes the new values into existingAnswer, so keep the previous file path
	previousFile := existingAnswer.File
	if err := database.DB.Model(&existingAnswer).Updates(updates).Error; err != nil {
		removeRecommendationAnswerFile(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	// the new file replaced the previous one
	if uploaded && previousFile != "" && previousFile != filePath {
		removeRecommendationAnswerFile(previousFile)
	}

	//fetch the updated answer
	var updatedAnswer models.RecommendationAnswer
	if err := database.DB.Preload("Recommendation.Tool").First(&updatedAnswer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}

// GET /recommendationAnswers/:id/file - Download the uploaded file of a recommendationAnswer
func DownloadRecommendationAnswerFile(c *gin.Context) {
	id, err := utils.ParseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return
	}

	var answer models.RecommendationAnswer
	if err := database.DB.First(&answer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	userId := c.GetString("user_id")
	// Validate journal ownership and existence
	if err := utils.CheckJournalAuthentication(answer.JournalID, userId); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if answer.File == "" || !isInRecommendationAnswerUploadDir(answer.File) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	if _, err := os.Stat(answer.File); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.FileAttachment(answer.File, filepath.Base(answer.File))
}
