package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	ctrl "server/controllers"
	"server/database"
	"server/models"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(
		sqlite.Open("file:test.db?mode=memory&cache=shared"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)

	db.AutoMigrate(
		&models.Lifecycle{},
		&models.Phase{},
		&models.Reflection{},
		&models.Recommendation{},
		&models.User{},
		&models.ReflectionAnswer{},
		&models.FurtherReflectionAnswer{},
		&models.RecommendationAnswer{},
	)

	return db
}

func TestGetLastUpdatedLifecycleItemForUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Log("Setting up test database...")
	db := setupTestDB()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	database.DB = db

	t.Log("Creating lifecycle...")
	lifecycle := models.Lifecycle{
		Title: "Lifecycle A",
	}
	db.Create(&lifecycle)

	t.Log("Creating phase...")
	phase := models.Phase{
		LifecycleID: lifecycle.ID,
		Lifecycle:   lifecycle,
	}
	db.Create(&phase)

	t.Log("Creating reflection...")
	reflection := models.Reflection{
		Title:   "Reflection X",
		PhaseID: phase.ID,
		Phase:   phase,
	}
	db.Create(&reflection)

	t.Log("Creating user and answers...")
	userUUID := uuid.New()
	userID := userUUID.String()

	reflectionAnswer := models.ReflectionAnswer{
		ReflectionID: reflection.ID,
		Reflection:   reflection,
		UserID:       userUUID,
		UpdatedAt:    time.Now(),
	}
	db.Create(&reflectionAnswer)

	recommendation := models.Recommendation{
		ReflectionID: reflection.ID,
		Reflection:   reflection,
	}
	db.Create(&recommendation)

	recommendationAnswer := models.RecommendationAnswer{
		RecommendationID: recommendation.ID,
		Recommendation:   recommendation,
		UserID:           userUUID,
		Form:             "test form",
		File:             "",
		CheckedDone:      false,
		UpdatedAt:        time.Now(),
	}
	db.Create(&recommendationAnswer)

	furtherReflectionAnswer := models.FurtherReflectionAnswer{
		ReflectionID: reflection.ID,
		Reflection:   reflection,
		UserID:       userUUID,
		Form:         "test further reflection form",
		UpdatedAt:    time.Now(),
	}
	db.Create(&furtherReflectionAnswer)

	t.Log("Setting up Gin router and endpoint...")
	r := gin.Default()

	r.GET("/progress/last-updated", func(c *gin.Context) {
		c.Set("user_id", userID)
		ctrl.GetLastUpdatedLifecycleItemForUser(c)
	})

	t.Log("Performing HTTP request...")
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		"/progress/last-updated",
		nil,
	)

	r.ServeHTTP(w, req)

	t.Logf("Response code: %d", w.Code)
	t.Logf("Response body: %s", w.Body.String())

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), "Reflection X") {
		t.Errorf("expected response to contain reflection title")
	}
}

func TestGetLastUpdatedLifecycleItemForMultipleUsers(t *testing.T) {
	db := setupTestDB()
	database.DB = db

	user1 := uuid.New()
	user2 := uuid.New()

	lifecycle := models.Lifecycle{Title: "L1"}
	db.Create(&lifecycle)
	phase := models.Phase{LifecycleID: lifecycle.ID, Lifecycle: lifecycle}
	db.Create(&phase)
	reflection := models.Reflection{Title: "R1", PhaseID: phase.ID, Phase: phase}
	db.Create(&reflection)

	db.Create(&models.ReflectionAnswer{ReflectionID: reflection.ID, Reflection: reflection, UserID: user1, UpdatedAt: time.Now()})
	db.Create(&models.ReflectionAnswer{ReflectionID: reflection.ID, Reflection: reflection, UserID: user2, UpdatedAt: time.Now()})

	r := gin.Default()
	r.GET("/progress/last-updated", func(c *gin.Context) {
		c.Set("user_id", user1.String())
		ctrl.GetLastUpdatedLifecycleItemForUser(c)
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/progress/last-updated", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	// Unmarshal response and check title
	type Response struct {
		Data struct {
			Title string `json:"title"`
		} `json:"data"`
	}
	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if response.Data.Title != "R1" {
		t.Errorf("expected title R1, got %s", response.Data.Title)
	}
}

func TestGetLastUpdatedLifecycleItemForUser_MissingUserID(t *testing.T) {
	db := setupTestDB()
	database.DB = db

	r := gin.Default()
	r.GET("/progress/last-updated", func(c *gin.Context) {
		// No user_id set
		ctrl.GetLastUpdatedLifecycleItemForUser(c)
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/progress/last-updated", nil)
	r.ServeHTTP(w, req)
	if w.Code == http.StatusOK {
		t.Errorf("expected error status when user_id missing, got 200")
	}
}

func TestGetLastUpdatedLifecycleItemForUser_InvalidUserID(t *testing.T) {
	db := setupTestDB()
	database.DB = db

	r := gin.Default()
	r.GET("/progress/last-updated", func(c *gin.Context) {
		c.Set("user_id", "not-a-uuid")
		ctrl.GetLastUpdatedLifecycleItemForUser(c)
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/progress/last-updated", nil)
	r.ServeHTTP(w, req)
	if w.Code == http.StatusOK {
		t.Errorf("expected error status for invalid user_id, got 200")
	}
}

func TestGetLastUpdatedLifecycleItemForUser_NoItems(t *testing.T) {
	db := setupTestDB()
	database.DB = db

	user := uuid.New()

	r := gin.Default()
	r.GET("/progress/last-updated", func(c *gin.Context) {
		c.Set("user_id", user.String())
		ctrl.GetLastUpdatedLifecycleItemForUser(c)
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/progress/last-updated", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "null") && w.Body.String() != "{}" {
		t.Errorf("expected empty or null response, got %s", w.Body.String())
	}
}

func TestGetLastUpdatedLifecycleItemForUser_ReflectionAnswer(t *testing.T) {
	db := setupTestDB()
	database.DB = db

	user := uuid.New()
	lifecycle := models.Lifecycle{Title: "L2"}
	db.Create(&lifecycle)
	phase := models.Phase{LifecycleID: lifecycle.ID, Lifecycle: lifecycle}
	db.Create(&phase)
	reflection := models.Reflection{Title: "ReflectionType", PhaseID: phase.ID, Phase: phase}
	db.Create(&reflection)
	db.Create(&models.ReflectionAnswer{ReflectionID: reflection.ID, Reflection: reflection, UserID: user, UpdatedAt: time.Now()})

	r := gin.Default()
	r.GET("/progress/last-updated", func(c *gin.Context) {
		c.Set("user_id", user.String())
		ctrl.GetLastUpdatedLifecycleItemForUser(c)
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/progress/last-updated", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "ReflectionType") {
		t.Errorf("expected response to contain ReflectionType")
	}
}

func TestGetLastUpdatedLifecycleItemForUser_RecommendationAnswer(t *testing.T) {
	db := setupTestDB()
	database.DB = db

	user := uuid.New()
	lifecycle := models.Lifecycle{Title: "L3"}
	db.Create(&lifecycle)
	phase := models.Phase{LifecycleID: lifecycle.ID, Lifecycle: lifecycle}
	db.Create(&phase)
	reflection := models.Reflection{Title: "RecType", PhaseID: phase.ID, Phase: phase}
	db.Create(&reflection)
	rec := models.Recommendation{ReflectionID: reflection.ID, Reflection: reflection}
	db.Create(&rec)
	db.Create(&models.RecommendationAnswer{RecommendationID: rec.ID, Recommendation: rec, UserID: user, Form: "form", UpdatedAt: time.Now()})

	r := gin.Default()
	r.GET("/progress/last-updated", func(c *gin.Context) {
		c.Set("user_id", user.String())
		ctrl.GetLastUpdatedLifecycleItemForUser(c)
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/progress/last-updated", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "RecType") {
		t.Errorf("expected response to contain RecType")
	}
}

func TestGetLastUpdatedLifecycleItemForUser_FurtherReflectionAnswer(t *testing.T) {
	db := setupTestDB()
	database.DB = db

	user := uuid.New()
	lifecycle := models.Lifecycle{Title: "L4"}
	db.Create(&lifecycle)
	phase := models.Phase{LifecycleID: lifecycle.ID, Lifecycle: lifecycle}
	db.Create(&phase)
	reflection := models.Reflection{Title: "FurtherType", PhaseID: phase.ID, Phase: phase}
	db.Create(&reflection)
	db.Create(&models.FurtherReflectionAnswer{ReflectionID: reflection.ID, Reflection: reflection, UserID: user, Form: "form", UpdatedAt: time.Now()})

	r := gin.Default()
	r.GET("/progress/last-updated", func(c *gin.Context) {
		c.Set("user_id", user.String())
		ctrl.GetLastUpdatedLifecycleItemForUser(c)
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/progress/last-updated", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "FurtherType") {
		t.Errorf("expected response to contain FurtherType")
	}
}

func TestGetLastUpdatedLifecycleItemForUser_DBError(t *testing.T) {
	db := setupTestDB()
	database.DB = db

	user := uuid.New()

	r := gin.Default()
	r.GET("/progress/last-updated", func(c *gin.Context) {
		c.Set("user_id", user.String())
		sqlDB, _ := db.DB()
		sqlDB.Close()
		ctrl.GetLastUpdatedLifecycleItemForUser(c)
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/progress/last-updated", nil)
	r.ServeHTTP(w, req)
	if w.Code == http.StatusOK {
		t.Errorf("expected error status for DB error, got 200")
	}
}
