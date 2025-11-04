package setup

import (
	"encoding/json"
	"os"
)

func loadFromFile(filename string, target interface{}) error {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, target)
}

func saveToFile(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}
