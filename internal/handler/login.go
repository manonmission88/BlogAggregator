package handler

import (
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/state"
)

// allow user login
func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("user login is required")
	}
	username := cmd.Args[0]
	// set the username
	err := s.Config.SetUser(username)
	if err != nil {
		return err
	}
	return nil
}
