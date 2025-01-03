package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func ReadConfig() (Config, error) {
	FilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(FilePath)
	if err != nil {
		return Config{}, fmt.Errorf("error finding file: %w", err)
	}

	ConfigStruct := Config{}
	if err := json.Unmarshal(data, &ConfigStruct); err != nil {
		return Config{}, err
	}
	return ConfigStruct, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	FilePath := filepath.Join(home, configFileName)
	return FilePath, nil
}

func write(cfg Config) error {
	FilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
