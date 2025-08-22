package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	// ConfigFileName is the name of the config file
	ConfigFileName = ".gemini-cli.json"
	// EnvKeyName is the environment variable name for the API key
	EnvKeyName = "GEMINI_API_KEY"
	// DefaultTimeout is the default timeout for API requests
	DefaultTimeout = 30 * time.Second
	// MaxRetries is the default number of retries for failed requests
	MaxRetries = 3
)

// Config holds the configuration for the CLI
type Config struct {
	APIKey      string        `json:"api_key,omitempty"`
	Timeout     time.Duration `json:"timeout,omitempty"`
	MaxRetries  int           `json:"max_retries,omitempty"`
	OutputStyle string        `json:"output_style,omitempty"`
}

// LoadConfig loads the configuration from file and environment
func LoadConfig() (*Config, error) {
	config := &Config{
		Timeout:    DefaultTimeout,
		MaxRetries: MaxRetries,
	}

	// Try to load from config file
	homeDir, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(homeDir, ConfigFileName)
		if data, err := os.ReadFile(configPath); err == nil {
			_ = json.Unmarshal(data, config)
		}
	}

	// Environment variables override file config
	if envKey := os.Getenv(EnvKeyName); envKey != "" {
		config.APIKey = envKey
	}

	return config, nil
}

// SaveConfig saves the configuration to file
func SaveConfig(config *Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	configPath := filepath.Join(homeDir, ConfigFileName)
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}

// ValidateFilePath checks if a file path is valid and accessible
func ValidateFilePath(path string, mustExist bool) error {
	if path == "" {
		return fmt.Errorf("file path cannot be empty")
	}

	// Clean the path to resolve any .. or . components
	cleanPath := filepath.Clean(path)

	// Check if path is absolute, if not make it absolute
	if !filepath.IsAbs(cleanPath) {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %v", err)
		}
		cleanPath = filepath.Join(cwd, cleanPath)
	}

	if mustExist {
		info, err := os.Stat(cleanPath)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("file does not exist: %s", path)
			}
			return fmt.Errorf("failed to access file: %v", err)
		}
		if info.IsDir() {
			return fmt.Errorf("path is a directory: %s", path)
		}
	} else {
		// Check if parent directory exists and is writable
		dir := filepath.Dir(cleanPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	return nil
}

// HTTPClientWithTimeout returns an HTTP client with timeout
func HTTPClientWithTimeout(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

// WithContext adds context to API requests
func WithContext(parent context.Context) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	ctx, _ := context.WithTimeout(parent, DefaultTimeout)
	return ctx
}
