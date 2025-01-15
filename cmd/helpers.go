package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadJSONFromFile(filePath string) (map[string]interface{}, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(bytes, &body); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return body, nil
}

func ReadJSONFromStdin() (map[string]interface{}, error) {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, err
	}

	return data, nil
}
