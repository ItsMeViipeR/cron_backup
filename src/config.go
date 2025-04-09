package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	BackupDir    string
	OutDir       string
	Delay        string
	Differential bool
}

func ReadConfigFile(file_path string) (*Config, error) {
	config := &Config{}

	// Open the JSON file
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON data into the config struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
