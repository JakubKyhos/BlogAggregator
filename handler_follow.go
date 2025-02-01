package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JakubKyhos/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerfollow(s *state, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("requires url")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find the feed: %s", err)
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

	fmt.Println("Following feed successfully:")
	fmt.Printf("%s\n", FeedFollow.FeedName)
	fmt.Printf("%s\n", FeedFollow.UserName)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerfollowing(s *state, cmd Command, user database.User) error {
	feedfollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}

	fmt.Printf("Feeds you're following:\n")
	fmt.Println("=====================================")

	if len(feedfollows) == 0 {
		fmt.Println("You're not following any feeds yet!")
		return nil
	}

	for i, feed := range feedfollows {
		fmt.Printf("%d. %s\n", i+1, feed.FeedName)
	}

	fmt.Println("=====================================")

	return nil
}

func handlerunfollow(s *state, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("requires url")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find the feed: %s", err)
	}
	var unfollowfeed = database.UnfollowFeedParams{
		Url:  feed.Url,
		Name: user.Name,
	}
	err = s.db.UnfollowFeed(context.Background(), unfollowfeed)
	if err != nil {
		return fmt.Errorf("couldn't find the feed: %s", err)
	}
	fmt.Printf("%s unfollowed.\n", feed.Name)
	return nil
}
