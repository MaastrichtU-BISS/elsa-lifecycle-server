// utils/auth.go
package utils

import (
	"errors"
	"fmt"
	"os"
	"server/database"
	"server/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const minJWTSecretLength = 32

var jwtKey []byte

// LoadJWTKey reads the token signing key from JWT_SECRET. Call it at startup, after loading .env
func LoadJWTKey() error {
	secret := os.Getenv("JWT_SECRET")
	if len(secret) < minJWTSecretLength {
		return fmt.Errorf("JWT_SECRET must be set to at least %d characters", minJWTSecretLength)
	}
	jwtKey = []byte(secret)
	return nil
}

func HashPassword(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	return string(bytes), err
}

func CheckPasswordHash(p, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(p)) == nil
}

func GenerateJWT(userID uuid.UUID) (string, error) {
	if len(jwtKey) == 0 {
		return "", errors.New("JWT key not loaded")
	}
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func CheckJournalAuthentication(journalID uint, userID string) error {
	// load the journal to check ownership and existence
	var journal models.Journal
	if err := database.DB.First(&journal, journalID).Error; err != nil {
		return err
	}

	// (Authentication) check if the journal belongs to the user
	if journal.UserID.String() != userID {
		return errors.New("you do not have access to this journal")
	}
	return nil
}
