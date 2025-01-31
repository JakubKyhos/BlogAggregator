package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/JakubKyhos/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("both name and url are required")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.configptr.CurrentUserName)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("user %s does not exist\n", s.configptr.CurrentUserName)
			return err
		}
		return err
	}

	var feedParams = database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("failed to create feed: %s", err)
	}

	var NewFeedFollow = database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	FeedFollow, err := s.db.CreateFeedFollow(context.Background(), NewFeedFollow)
	if err != nil {
		return fmt.Errorf("failed to follow feed: %s", err)
	}

	fmt.Println("Feed created and followed successfully:")
	printFeed(feed, user, FeedFollow)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeed(feed database.Feed, user database.User, feedfollow database.CreateFeedFollowRow) {
	fmt.Printf(" * Name:    		%v\n", feed.Name)
	fmt.Printf(" * URL:    			%v\n", feed.Url)
	fmt.Printf(" * User:    		%v\n", user.Name)
	fmt.Printf(" * Now following feed: %s as user: %s\n", feedfollow.FeedName, feedfollow.UserName)
}
