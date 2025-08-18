package utils

import (
	"fmt"
	"strconv"
	"time"
)

type invalidFormatError struct {
	timeStr string
}

func (e invalidFormatError) Error() string {
	return fmt.Sprintf("invalid time format: %s (expected RFC3339 or Unix timestamp)", e.timeStr)
}

func InvalidFormatError(timeStr string) error {
	return invalidFormatError{timeStr}
}

var _ error = invalidFormatError{}

// ParseTime parses a time string in either RFC3339 or Unix timestamp format.
func ParseTime(timeStr string) (time.Time, error) {
	// Try RFC3339 format first
	if t, err := time.Parse(time.RFC3339, timeStr); err == nil {
		return t, nil
	}

	// Try Unix timestamp format
	if unix, err := strconv.ParseInt(timeStr, 10, 64); err == nil {
		return time.Unix(unix, 0), nil
	}

	return time.Time{}, InvalidFormatError(timeStr)
}
