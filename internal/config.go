package internal

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func ReadConfig(path string) (any, error) {
	var config any

	// Open YAML file
	file, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	// Decode YAML file to struct
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			log.Println(err.Error())
		}
	}

	return config, nil
}
