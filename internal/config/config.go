package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homeDir, configFileName)
	return path, nil
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(content, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(*c)
	
	if err != nil {
		return err
	}

	return nil
}