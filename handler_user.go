package main

import (
	"context"
	"fmt"
	"os"

	"database/sql"
)

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username required")
	}

	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("user %s does not exist\n", username)
			os.Exit(1)
		}
		return err
	}

	s.configptr.CurrentUserName = username
	err = s.configptr.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("user set to: %s\n", username)
	return nil
}
