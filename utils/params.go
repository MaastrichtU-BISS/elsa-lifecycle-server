package utils

import (
	"errors"
	"strconv"
)

var ErrInvalidID = errors.New("invalid id")

// ParseID parses a positive integer ID from a URL parameter or query value.
// Always parse IDs before passing them to GORM: a raw string passed to First/Find
// as a condition is inserted into the query as SQL.
func ParseID(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil || id == 0 {
		return 0, ErrInvalidID
	}
	return uint(id), nil
}
