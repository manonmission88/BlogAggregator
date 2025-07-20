package handler

import (
	"context"
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// add followers to the feed
func HandlerUnfollow(s *state.State, cmd Command, user database.User) error {
	args := cmd.Args
	if len(args) < 1 {
		return fmt.Errorf("usage : unfollow <feed_url>")
	}
	feedUrl := args[0]
	err := s.DbQueries.UnfollowFeed(context.Background(), database.UnfollowFeedParams{UserID: user.ID, Url: feedUrl})
	if err != nil {
		return fmt.Errorf("couldnot delete : %w", err)
	}
	fmt.Printf("%s unfollowed feed %s\n", user.Name, feedUrl)
	return nil
}
