package common

import "fmt"

// Common error types for consistent error handling
var (
	ErrAPIKeyNotSet    = fmt.Errorf("API key not set. Use 'gemini-cli config -k YOUR_API_KEY' to set it")
	ErrInvalidFilePath = fmt.Errorf("invalid file path provided")
	ErrAPIRequest      = fmt.Errorf("API request failed")
	ErrFileAccess      = fmt.Errorf("file access error")
)

// APIError represents an error from the Gemini API
type APIError struct {
	StatusCode int
	Message    string
	RetryCount int
}

func (e *APIError) Error() string {
	if e.RetryCount > 0 {
		return fmt.Sprintf("API request failed after %d retries with status %d: %s", e.RetryCount, e.StatusCode, e.Message)
	}
	return fmt.Sprintf("API request failed with status %d: %s", e.StatusCode, e.Message)
}

// IsRetryableError determines if an error should trigger a retry
func IsRetryableError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode >= 500 || apiErr.StatusCode == 429 // Retry on server errors and rate limits
	}
	return false
}
