package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(*c)
	
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	content, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer content.Close()

	decoder := json.NewDecoder(content)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}


	return cfg, nil
}

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

	jsonData, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer jsonData.Close()

	encoder := json.NewEncoder(jsonData)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}



