package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config  struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const filename = ".gatorconfig.json"

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("error getting config file path: %v", err)
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file path: %v", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling data: %v", err)
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(*c)
	if err != nil {
		return fmt.Errorf("error writing config to a file: %v", err)
	}
	return nil
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %v", err)
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error marshalling config: %v", err)
	}
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error accessing home directory: %v", err)
	}
	return fmt.Sprintf("%s%c%s", homedir, os.PathSeparator, filename), nil
}