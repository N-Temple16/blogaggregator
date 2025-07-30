package main

import (
	"blogaggregator/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName

	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if len(cmd.StringSlice) != 2 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	
	name := cmd.StringSlice[0]
	url := cmd.StringSlice[1]	

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams {
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to auto-follow feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsWithUser(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds list: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
		fmt.Printf("%s\n", feed.Url)
		fmt.Printf("%s\n", feed.UserName)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * URL:    %v\n", feed.Url)
	fmt.Printf(" * UserID:    %v\n", feed.UserID)
}