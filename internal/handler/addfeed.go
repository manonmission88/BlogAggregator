package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// allow register user to the database
func HandlerAddFeed(s *state.State, cmd Command) error {
	args := cmd.Args
	if len(args) < 2 {
		return fmt.Errorf("usage addfeed <name of the feed> < url feed> Example : addfeed <Hacker News RSS> <https://hnrss.org/newest> ")
	}
	name := args[0]
	url := args[1]

	userName, err := s.DbQueries.GetUserByName(context.Background(), s.Config.CurrentUser)
	if err != nil {
		return fmt.Errorf("could not get current user")
	}

	_, err = s.DbQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    userName.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed")
	}
	return nil
}
