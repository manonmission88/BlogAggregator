package handler

import (
	"context"
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// Ensures the user is logged in before proceeding.
// Returns an error if no user is authenticated.
func MiddlewareLoggedIn(handler func(s *state.State, cmd Command, user database.User) error) func(*state.State, Command) error {
	return func(s *state.State, cmd Command) error {
		user, err := s.DbQueries.GetUserByName(context.Background(), s.Config.CurrentUser)
		if err != nil {
			return fmt.Errorf("user is not logged in %w\n", err)
		}
		return handler(s, cmd, user)
	}
}
