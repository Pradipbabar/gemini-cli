package common

import (
	"os"
	"path/filepath"
)

// WriteFile writes data to a file, creating parent directories if needed
func WriteFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
