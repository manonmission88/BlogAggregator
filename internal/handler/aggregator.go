package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// allow register user to the database
func HandlerAgg(s *state.State, cmd Command, user database.User) error {
	args := cmd.Args
	if len(args) < 1 {
		return fmt.Errorf("usage agg <time lapse>")
	}
	timeRequest, err := time.ParseDuration(args[0])
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Printf("Collecting feed in time ---> %s\n", timeRequest)

	// create time ticker -> sends signal on every that time gap
	ticker := time.NewTicker(timeRequest)
	// defer ticker.Stop()

	for ; ; <-ticker.C {
		ScrapeFeeds(s)
	}
}
