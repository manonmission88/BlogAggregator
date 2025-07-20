package handler

import (
	"context"
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/rss"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// scrape all the feeds based on the last updated
func ScrapeFeeds(s *state.State) error {
	// fetch the next feed
	nextFeed, err := s.DbQueries.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("could not fetch the feed: %w", err)
	}

	fmt.Printf("Fetching feed URL: %s\n", nextFeed.Url)

	feed, err := s.DbQueries.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("could not mark feed fetched: %w", err)
	}

	feedDetail, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("could not fetch feed %s: %w", nextFeed.Url, err)
	}
	fmt.Printf("Feed Title: %s | Items count: %d\n", feedDetail.Channel.Title, len(feedDetail.Channel.Item))
	for _, fdetail := range feedDetail.Channel.Item {
		fmt.Printf("Title : %s\n", fdetail.Title)
	}

	return nil
}
