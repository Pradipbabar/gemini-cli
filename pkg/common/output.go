package common

import (
	"fmt"
	"strings"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[36m"
)

// Output formats
const (
	OutputStylePlain = "plain"
	OutputStyleJSON  = "json"
	OutputStyleColor = "color"
)

// FormatError returns a formatted error message
func FormatError(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%s[Error]%s %v", ColorRed, ColorReset, err)
}

// FormatSuccess returns a formatted success message
func FormatSuccess(msg string) string {
	return fmt.Sprintf("%s[Success]%s %s", ColorGreen, ColorReset, msg)
}

// FormatInfo returns a formatted info message
func FormatInfo(msg string) string {
	return fmt.Sprintf("%s[Info]%s %s", ColorBlue, ColorReset, msg)
}

// FormatWarning returns a formatted warning message
func FormatWarning(msg string) string {
	return fmt.Sprintf("%s[Warning]%s %s", ColorYellow, ColorReset, msg)
}

// FormatJSON formats data as JSON with proper indentation
func FormatJSON(key string, value interface{}) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("{\n  %q: %q\n}", key, v)
	default:
		return fmt.Sprintf("{\n  %q: %v\n}", key, v)
	}
}

// FormatProgress returns a progress message with spinner
func FormatProgress(msg string) string {
	return fmt.Sprintf("%s[*]%s %s", ColorBlue, ColorReset, msg)
}

// PromptForInput prompts the user for input with optional default value
func PromptForInput(prompt string, defaultValue string) string {
	if defaultValue != "" {
		fmt.Printf("%s (%s): ", prompt, defaultValue)
	} else {
		fmt.Printf("%s: ", prompt)
	}

	var input string
	fmt.Scanln(&input)

	if input == "" {
		return defaultValue
	}
	return strings.TrimSpace(input)
}
