package pkg

import (
	"fmt"
	"os"
)

func ReadFromFile(filePath string) (string, error) {
	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(content), nil
}

func ReadFromFileandSave(filePath, savepath string) (string, error) {
	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}
	fmt.Println(savepath)

	return string(content), nil
}
