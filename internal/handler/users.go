package handler

import (
	"context"
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/state"
)

// allow register user to the database
func HandlerUsers(s *state.State, cmd Command) error {

	names, err := s.DbQueries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("no users found")
	}
	if len(names) == 0 {
		return fmt.Errorf("No any users on the database")
	}
	current_user := s.Config.CurrentUser
	for _, name := range names {
		if name == current_user {
			fmt.Printf("* %s (current)\n", name)
		} else {
			fmt.Printf("* %s \n", name)
		}
	}
	return nil
}
