package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read(filename string) (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("Error locating home directory", err)
	}

	configFileData, err := os.ReadFile(home + "/" + filename)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading gatorconfig file", err)
	}

	gatorConfig := Config{}

	err = json.Unmarshal(configFileData, &gatorConfig)
	if err != nil {
		return Config{}, fmt.Errorf("Error parsing json", err)
	}

	return gatorConfig, nil
}
