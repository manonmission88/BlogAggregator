package handler

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// allow register user to the database
func HandlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("name is required")
	}
	userName := cmd.Args[0]
	// check if the username already exists
	_, err := s.DbQueries.GetUserByName(context.Background(), userName)
	if err == nil {
		os.Exit(1)
	}
	// set the username
	_, err = s.DbQueries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}) //
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	err = s.Config.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Printf("User %s was created\n", userName)
	return nil
}
