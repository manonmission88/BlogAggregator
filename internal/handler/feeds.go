package handler

import (
	"context"
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/state"
)

// return all the feeds in the database
func HandlerFeed(s *state.State, cmd Command) error {

	feedsData, err := s.DbQueries.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not get the feeds : DB error")
	}

	for _, feed := range feedsData {
		fmt.Printf("Name of the feed : %s\n", feed.FeedName)
		fmt.Printf("Url of the feed : %s\n", feed.UrlName)
		fmt.Printf("User who created the feed : %s\n", feed.UserName)
		fmt.Println()
	}
	return nil
}
