package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/manonmission88/BlogAggregator/internal/state"
)

// allow user login
func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("user login is required")
	}
	userName := cmd.Args[0]
	// check if the username doesnot exists
	_, err := s.DbQueries.GetUserByName(context.Background(), userName)
	if err != nil {
		os.Exit(1)
	}
	// set the username
	err = s.Config.SetUser(userName)
	if err != nil {
		return err
	}

	return nil
}
