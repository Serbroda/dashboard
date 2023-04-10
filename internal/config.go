package internal

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Title    string    `json:"title"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

type Item struct {
	Name                             string `json:"name"`
	Url                              string `json:"url"`
	CheckAvailability                bool   `json:"checkAvailability"`
	CheckAvailabilityIntervalSeconds int32  `json:"checkAvailabilityIntervalSeconds"`
}

func ReadConfig(path string) (any, error) {
	var config Config

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &config)
	return config, nil
}
