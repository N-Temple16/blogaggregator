package main

import (
	"blogaggregator/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName

	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if len(cmd.StringSlice) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	
	url := cmd.StringSlice[0]

	feed, err := s.db.GetFeedWithUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to get specified feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams {
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	fmt.Printf(" * Feed:    %v\n", feedFollow.FeedName)
	fmt.Printf(" * User:    %v\n", feedFollow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("failed to get feed follows for the given user: %w", err)
	}

	for _, feed := range feedFollows {
		fmt.Printf("%s\n", feed.FeedName)
	}

	return nil
}