package pkg

import (
	"fmt"
	"os"
)

func SetEnvironmentVariable(key, value string) error {
	err := os.Setenv(key, value)
	if err != nil {
		return fmt.Errorf("failed to set environment variable: %v", err)
	}
	return nil
}
