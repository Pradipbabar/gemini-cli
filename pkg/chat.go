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

type RequestBody struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"Parts"`
	} `json:"contents"`
}

type ResponseBody struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
		FinishReason string `json:"finishReason"`
		Index        int    `json:"index"`
		SafetyRatings []struct {
			Category    string `json:"category"`
			Probability string `json:"probability"`
		} `json:"safetyRatings"`
	} `json:"candidates"`
	PromptFeedback struct {
		SafetyRatings []struct {
			Category    string `json:"category"`
			Probability string `json:"probability"`
		} `json:"safetyRatings"`
	} `json:"promptFeedback"`
}

func GenerateContent(inputText string) (string, error) {
	apikey := os .Getenv("GEMINI_API_KEY")
	if apikey == "" {
		fmt.Println(apikey)
		return "", fmt.Errorf("Api Key not set")
	}
	// Prepare the request body
	requestBody := RequestBody{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"Parts"`
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

	// Convert the request body to JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Make the API request
	resp, err := http.Post(fmt.Sprintf(gptAPIURL,apikey), "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the response JSON
	var response ResponseBody
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	// Extract and return the generated text
	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		return response.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("no generated text found in response")
}

func GenerateContentAndSave(inputText, filePath string) error {
	// Prepare the request body
	requestBody := RequestBody{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"Parts"`
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

	// Convert the request body to JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Make the API request
	resp, err := http.Post(gptAPIURL, "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the response JSON
	var response ResponseBody
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
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
