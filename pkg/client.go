package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Pradipbabar/gimini-cli/pkg/common"
)

// GeminiClient represents a client for the Gemini API
type GeminiClient struct {
	config *common.Config
	client *http.Client
}

// NewGeminiClient creates a new Gemini API client
func NewGeminiClient(config *common.Config) *GeminiClient {
	return &GeminiClient{
		config: config,
		client: common.HTTPClientWithTimeout(config.Timeout),
	}
}

// GenerateContent generates content from the given input text
func (c *GeminiClient) GenerateContent(ctx context.Context, inputText string) (string, error) {
	if c.config.APIKey == "" {
		return "", common.ErrAPIKeyNotSet
	}

	requestBody := RequestBody{
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

	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	var response *http.Response
	var lastErr error

	for i := 0; i <= c.config.MaxRetries; i++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, gptAPIURL, bytes.NewBuffer(requestBytes))
		if err != nil {
			return "", fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+c.config.APIKey)

		response, err = c.client.Do(req)
		if err == nil {
			break
		}

		lastErr = err
		if i < c.config.MaxRetries && common.IsRetryableError(err) {
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		break
	}

	if response == nil {
		return "", &common.APIError{
			Message:    lastErr.Error(),
			RetryCount: c.config.MaxRetries,
		}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return "", &common.APIError{
			StatusCode: response.StatusCode,
			Message:    string(body),
		}
	}

	var responseBody ResponseBody
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(responseBody.Candidates) > 0 && len(responseBody.Candidates[0].Content.Parts) > 0 {
		return responseBody.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("no content generated in response")
}

// GenerateContentToFile generates content and saves it to a file
func (c *GeminiClient) GenerateContentToFile(ctx context.Context, inputText, filePath string) error {
	if err := common.ValidateFilePath(filePath, false); err != nil {
		return fmt.Errorf("invalid output file: %w", err)
	}

	content, err := c.GenerateContent(ctx, inputText)
	if err != nil {
		return err
	}

	return common.WriteFile(filePath, []byte(content))
}
