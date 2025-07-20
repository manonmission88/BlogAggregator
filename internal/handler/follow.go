package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// add followers to the feed
func HandlerFollow(s *state.State, cmd Command, user database.User) error {
	args := cmd.Args
	if len(args) < 1 {
		return fmt.Errorf("usage follow < url feed> Example : follow  <https://hnrss.org/newest> : ")
	}
	url := args[0]
	feedId, err := s.DbQueries.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	_, err = s.DbQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feedId.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create followers feed: %w", err)
	}
	fmt.Printf("feed name : %s\n", feedId.Name)
	fmt.Printf("user name : %s\n", user.Name)

	return nil
}
