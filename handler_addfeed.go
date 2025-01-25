package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/JakubKyhos/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *state, cmd Command) error {
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

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf(" * ID:      		%v\n", feed.ID)
	fmt.Printf(" * CreatedAt:    	%v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt:    	%v\n", feed.UpdatedAt)
	fmt.Printf(" * Name:    		%v\n", feed.Name)
	fmt.Printf(" * URL:    			%v\n", feed.Url)
	fmt.Printf(" * User:    		%v\n", user.Name)
}
