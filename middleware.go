package main

import (
	"context"

	"github.com/JakubKyhos/blogaggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd Command, user database.User) error) func(*state, Command) error {
	return func(s *state, cmd Command) error {
		user, err := s.db.GetUser(context.Background(), s.configptr.CurrentUserName)
		if err != nil {
			return err
		}
		err = handler(s, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
}
