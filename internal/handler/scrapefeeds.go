package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/manonmission88/BlogAggregator/internal/database"
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

	for _, fd := range feedDetail.Channel.Item {
		// Parse fd.PubDate string to time.Time
		var pubTime time.Time
		var errPubDate error
		if fd.PubDate != "" {
			pubTime, errPubDate = time.Parse(time.RFC1123Z, fd.PubDate)
			if errPubDate != nil {
				log.Printf("Failed to parse PubDate '%s': %v\n", fd.PubDate, errPubDate)
			}
		}
		_, err := s.DbQueries.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     fd.Title,
			Url:       fd.Link,
			Description: sql.NullString{
				String: fd.Description,
				Valid:  fd.Description != "",
			},
			PublishedAt: sql.NullTime{
				Time:  pubTime,
				Valid: fd.PubDate != "" && errPubDate == nil,
			},
			FeedID: feed.ID,
		})

		if err != nil {
			continue
		}
	}
	return nil
}
