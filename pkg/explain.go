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

	data, err := GenerateContent(string(content))
			if err != nil {
				return "", err
			} else{
				return data, nil
			}
}

func ReadFromFileandSave(filePath, savepath string) (string, error) {
	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}
	
	err = GenerateContentAndSave(string(content),savepath)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("%v",err)
	}
	return string(content), nil
}
