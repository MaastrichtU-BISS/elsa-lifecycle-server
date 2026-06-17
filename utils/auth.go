// utils/auth.go
package utils

import (
	"errors"
	"server/database"
	"server/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key") // TODO: Replace with your secret key fom environment variable

func HashPassword(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	return string(bytes), err
}

func CheckPasswordHash(p, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(p)) == nil
}

func GenerateJWT(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func CheckJournalAuthentication(journalID string, userID string) error {
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
