package handler

import (
	"context"
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/rss"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

const url = "https://www.wagslane.dev/index.xml"

// allow register user to the database
func HandlerAgg(s *state.State, cmd Command) error {
	feedData, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldnot fetch the data")
	}
	// print all the struct field
	fmt.Printf("Feed Title: %s\n", feedData.Channel.Title)
	fmt.Printf("Feed Description: %s\n", feedData.Channel.Description)
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Item Title: %s\n", item.Title)
		fmt.Printf("Item Link: %s\n", item.Link)
		fmt.Printf("Item Description: %s\n", item.Description)
		fmt.Println("---")
	}
	return nil

}
