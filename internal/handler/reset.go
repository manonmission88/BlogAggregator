package handler

import (
	"context"
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// allow register user to the database
func HandlerReset(s *state.State, cmd Command, user database.User) error {
	err := s.DbQueries.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Println("something is wrong with the db")
	}
	fmt.Println("Db table is succefully reset")
	return nil
}
