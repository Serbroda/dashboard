package internal

import (
	"encoding/json"
	"os"
)

func ReadConfig[T any](path string) (T, error) {
	var config T

	content, err := ReadFileBytes(path)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func SaveConfig[T any](path string, config T) error {
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}
