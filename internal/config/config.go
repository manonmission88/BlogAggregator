package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFile = ".gatorconfig.json"

// according to gatorconfig
type config struct {
	DbUrl       string `json:"db_url"`
	CurrentUser string `json:"current_user"`
}

// get the json file path
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFile), nil
}

// read the json file
func read() (config, error) {
	var cfg config // empty config struct
	path, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}
	// read file
	file, err := os.ReadFile(path)
	if err != nil {
		return cfg, nil
	}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return cfg, nil
	}
	return cfg, nil

}

// set the user name
func (cfg *config) SetUser(username string) error {
	cfg.CurrentUser = username
	return write(*cfg) // write in the config struct
}

// write on the config
func write(cfg config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644) // 0644 -read/write access

}
