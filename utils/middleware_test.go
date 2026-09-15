package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func serveWithToken(token string) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/", AuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, c.GetString("user_id"))
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestLoadJWTKeyRequiresLongSecret(t *testing.T) {
	t.Setenv("JWT_SECRET", "")
	if err := LoadJWTKey(); err == nil {
		t.Fatal("expected error for missing JWT_SECRET")
	}

	t.Setenv("JWT_SECRET", "too-short")
	if err := LoadJWTKey(); err == nil {
		t.Fatal("expected error for short JWT_SECRET")
	}
}

func TestAuthMiddleware(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret-that-is-at-least-32-characters")
	if err := LoadJWTKey(); err != nil {
		t.Fatal(err)
	}

	userID := uuid.New()
	valid, err := GenerateJWT(userID)
	if err != nil {
		t.Fatal(err)
	}
	if w := serveWithToken(valid); w.Code != http.StatusOK || w.Body.String() != userID.String() {
		t.Fatalf("valid token: got %d %q", w.Code, w.Body.String())
	}

	claims := jwt.MapClaims{"user_id": userID.String(), "exp": time.Now().Add(time.Hour).Unix()}

	wrongKey, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("my_secret_key"))
	if w := serveWithToken(wrongKey); w.Code != http.StatusUnauthorized {
		t.Fatalf("token signed with another key: got %d, want 401", w.Code)
	}

	unsigned, _ := jwt.NewWithClaims(jwt.SigningMethodNone, claims).SignedString(jwt.UnsafeAllowNoneSignatureType)
	if w := serveWithToken(unsigned); w.Code != http.StatusUnauthorized {
		t.Fatalf("unsigned token: got %d, want 401", w.Code)
	}

	noUser, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"exp": time.Now().Add(time.Hour).Unix()}).SignedString(jwtKey)
	if w := serveWithToken(noUser); w.Code != http.StatusUnauthorized {
		t.Fatalf("token without user_id: got %d, want 401", w.Code)
	}
}
