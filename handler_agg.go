package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAggregator(s *state, cmd command) error {
	time_between_reqs, err := time.ParseDuration(cmd.StringSlice[0])
	if err != nil {
		fmt.Println("Error parsing duration:", err)
	}
	ticker := time.NewTicker(time_between_reqs)

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch the next feed: %w", err)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark fetched feed: %w", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	for _, f := range feed.Channel.Item {
		fmt.Printf("* %s\n", f.Title)
	}

	return nil
}