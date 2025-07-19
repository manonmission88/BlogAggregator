package handler

import (
	"context"
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/state"
)

// add followers to the feed
func HandlerFollowing(s *state.State, cmd Command) error {

	user, err := s.DbQueries.GetUserByName(context.Background(), s.Config.CurrentUser)
	if err != nil {
		return err
	}
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
