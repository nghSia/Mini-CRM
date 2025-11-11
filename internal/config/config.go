package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type InputTarget struct {
	Id    int    `json: "Id"`
	Name  string `json: "Name"`
	Email string `json: "Email"`
}

func LoadTargetsFromFile(input string) ([]InputTarget, error) {
	data, err := os.ReadFile(input)
	if err != nil {
		return nil, fmt.Errorf("Error while trying to read data %w", err)
	}

	var targets []InputTarget
	if err = json.Unmarshal(data, &targets); err != nil {
		return nil, fmt.Errorf("Error during deserialization %s : %w", input, err)
	}
	return targets, nil
}

func SaveTargetsToFile(targets []InputTarget, output string) error {
	data, err := json.MarshalIndent(targets, "", "  ")
	if err != nil {
		return fmt.Errorf("Error during serialization %s : %w", output, err)
	}

	if err := os.WriteFile(output, data, 0644); err != nil {
		return fmt.Errorf("Error while trying to write data %s : %w", output, err)
	}
	return nil
}
