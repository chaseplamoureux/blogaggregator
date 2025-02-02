package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const ConfigFileName = ".gatorconfig.json"

func getConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(home, ConfigFileName)
	return path

}

func Read() (Config, error) {
	path := getConfigFilePath()
	configFileData, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading gatorconfig file %v", err)
	}

	gatorConfig := Config{}

	err = json.Unmarshal(configFileData, &gatorConfig)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing json %v", err)
	}

	return gatorConfig, nil
}

func (conf *Config) SetUser(username string) {
	conf.Username = username
	err := write(*conf)
	if err != nil {
		fmt.Println(err)
	}
}

func write(conf Config) error {
	data, err := json.Marshal(conf)
	if err != nil {
		return fmt.Errorf("error parsing config to json %v", err)
	}

	path := getConfigFilePath()
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file %v", err)
	}
	return nil
}
