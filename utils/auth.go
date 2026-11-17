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

// bcryptCost is the work factor used when hashing passwords.
const bcryptCost = 14

// MaxPasswordBytes is bcrypt's upper limit; longer inputs cannot be hashed safely
// (bcrypt silently truncates beyond this point).
const MaxPasswordBytes = 72

// ErrPasswordTooLong is returned when a password exceeds bcrypt's input limit.
var ErrPasswordTooLong = fmt.Errorf("password must be at most %d bytes", MaxPasswordBytes)

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

// ValidatePassword checks a plaintext password against bcrypt's input constraints.
func ValidatePassword(p string) error {
	if p == "" {
		return errors.New("password is required")
	}
	if len(p) > MaxPasswordBytes {
		return ErrPasswordTooLong
	}
	return nil
}

func HashPassword(p string) (string, error) {
	if err := ValidatePassword(p); err != nil {
		return "", err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(p), bcryptCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
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
