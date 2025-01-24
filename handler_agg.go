package main

import (
	"context"
	"fmt"

	rssfeed "github.com/JakubKyhos/blogaggregator/internal/rssFeed"
)

func handleAggCmd(s *state, cmd Command) error {
	ctx := context.Background()
	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := rssfeed.FetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}

	fmt.Printf("Feed Title: %s\n", feed.Channel.Title)
	fmt.Printf("Feed Description: %s\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("Item Title: %s\n", item.Title)
		fmt.Printf("Item Description: %s\n", item.Description)
		fmt.Printf("Item Link: %s\n", item.Link)
		fmt.Printf("Publication Date: %s\n", item.PubDate)
	}
	return nil
}
