package state

import "github.com/manonmission88/BlogAggregator/internal/config"

// State holds a pointer to the Config struct.
type State struct {
	Config *config.Config
}
