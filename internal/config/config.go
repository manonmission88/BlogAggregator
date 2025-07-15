package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFile = ".gatorconfig.json"

// Exported Struct and Fields
type Config struct {
	DbUrl       string `json:"db_url"`
	CurrentUser string `json:"current_user"`
}

// Returns full path like /home/user/.gatorconfig.json
func getConfigFilePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, configFile), nil
}

// Read config from file
func Read() (*Config, error) {
	var cfg Config
	path, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Set the current user and save
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUser = username
	return write(cfg)
}

// Write config back to file
func write(cfg *Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
