package handler

import (
	"context"
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// add followers to the feed
func HandlerFollowing(s *state.State, cmd Command, user database.User) error {

	followersData, err := s.DbQueries.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, follower := range followersData {
		fmt.Println(follower.FeedName)
		fmt.Println(follower.UserName)
		fmt.Println()
	}
	return nil
}
