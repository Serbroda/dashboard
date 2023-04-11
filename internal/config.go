package internal

import (
	"encoding/json"
	"log"
	"os"
)

func ReadConfig[T any](path string) (T, error) {
	var config T

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &config)
	return config, nil
}
