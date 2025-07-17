package state

import "main/internal/config"

// State holds a pointer to the Config struct.
type State struct {
	Config *config.Config
}
