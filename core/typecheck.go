package core

import (
	"strconv"
	"time"
)

// ValidationError holds key and error reason
type ValidationError struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

// ValidateTypes checks if certain env vars match expected types
func ValidateTypes(env map[string]string) []ValidationError {
	var errors []ValidationError

	// PORT must be an integer
	if val, ok := env["PORT"]; ok {
		if _, err := strconv.Atoi(val); err != nil {
			errors = append(errors, ValidationError{
				Key:   "PORT",
				Error: "must be an integer",
			})
		}
	}

	// DEBUG must be a boolean
	if val, ok := env["DEBUG"]; ok {
		if val != "true" && val != "false" {
			errors = append(errors, ValidationError{
				Key:   "DEBUG",
				Error: "must be 'true' or 'false'",
			})
		}
	}

	// TIMEOUT must be a valid duration (e.g., 5s, 1m)
	if val, ok := env["TIMEOUT"]; ok {
		if _, err := time.ParseDuration(val); err != nil {
			errors = append(errors, ValidationError{
				Key:   "TIMEOUT",
				Error: "must be a valid duration (e.g. 5s, 1m)",
			})
		}
	}

	return errors
}
