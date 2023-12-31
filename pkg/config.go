package pkg

import (
	"fmt"
	"os"
)

func SetEnvironmentVariable( value string) error {
	key := "GEMINI_API_KEY"
	err := os.Setenv(key, value)
	if err != nil {
		return fmt.Errorf("failed to set environment variable: %v", err)
	}
	
	return nil
}
