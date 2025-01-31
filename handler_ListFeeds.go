package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd Command) error {
	feed, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}
	for _, feed := range feed {
		fmt.Printf("%v\n", feed.Name)
		fmt.Printf("%v\n", feed.Url)
		fmt.Printf("%v\n", feed.Name_2)
	}

	return nil
}
