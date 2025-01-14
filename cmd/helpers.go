package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadJSON(filePath string) (map[string]interface{}, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file: %w", err)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(bytes, &body); err != nil {
		return nil, fmt.Errorf("Failed to parse JSON: %w", err)
	}

	return body, nil
}
