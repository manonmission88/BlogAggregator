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
func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

// Set the current user and save
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUser = username
	return write(*cfg)
}

// Write config back to file
func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	// Debug: print config before writing
	file, err := os.Create(fullPath) // truncate the path if its already exist
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(cfg); err != nil {
		return err
	}

	return nil
}
