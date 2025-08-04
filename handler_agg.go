package main

import (
	"blogaggregator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
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

	for _, item := range feed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams {
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

	return nil
}