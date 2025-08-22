package pkg

import (
	"fmt"
	"os"
)

// ReadFromFileandSave reads a file, generates content, saves it, and returns the original content or error
func ReadFromFileandSave(filePath, savePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}
	err = GenerateContentAndSave(string(content), savePath)
	if err != nil {
		return "", fmt.Errorf("failed to generate or save content: %v", err)
	}
	return string(content), nil
}

// SafeFileStat checks if a file exists and is not a directory, returns error if not accessible
func SafeFileStat(filePath string) (os.FileInfo, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("path is a directory, not a file")
	}
	return info, nil
}

func ReadFromFile(filePath string) (string, error) {
	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	data, err := GenerateContent(string(content))
	if err != nil {
		return "", err
	}
	return data, nil
}
