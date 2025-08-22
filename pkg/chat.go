package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const gptAPIURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent"

func createRequestBody(inputText string) RequestBody {
	return RequestBody{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{
						Text: inputText,
					},
				},
			},
		},
	}
}

type APIError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

func makeAPIRequest(requestBody RequestBody) (*ResponseBody, error) {
	// Convert the request body to JSON and log it
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}
	fmt.Printf("Request body: %s\n", string(requestBodyBytes))

	// Create request
	req, err := http.NewRequest("POST", gptAPIURL+"?key="+os.Getenv("GEMINI_API_KEY"), bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Make the API request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check if response contains an error
	if resp.StatusCode != http.StatusOK {
		var apiError APIError
		if err := json.Unmarshal(respBody, &apiError); err == nil {
			return nil, fmt.Errorf("API error: %s", apiError.Error.Message)
		}
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse the response JSON
	var response ResponseBody
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return &response, nil
}

func GenerateContent(inputText string) (string, error) {
	apikey := os.Getenv("GEMINI_API_KEY")
	if apikey == "" {
		return "", fmt.Errorf("api key not set")
	}

	requestBody := createRequestBody(inputText)
	response, err := makeAPIRequest(requestBody)
	if err != nil {
		return "", err
	}

	// Extract and return the generated text
	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		return response.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("no generated text found in response")
}

func GenerateContentAndSave(inputText, filePath string) error {
	apikey := os.Getenv("GEMINI_API_KEY")
	if apikey == "" {
		return fmt.Errorf("api key not set")
	}

	requestBody := createRequestBody(inputText)
	response, err := makeAPIRequest(requestBody)
	if err != nil {
		return err
	}

	// Extract and save the generated text to the file
	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		generatedText := response.Candidates[0].Content.Parts[0].Text
		err := os.WriteFile(filePath, []byte(generatedText), 0644)
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
		return nil
	}

	return fmt.Errorf("no generated text found in response")
}
