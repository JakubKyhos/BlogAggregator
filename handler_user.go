package main

import (
	"fmt"
)

func HandlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username required")
	}

	username := cmd.Args[0]
	s.Configptr.CurrentUserName = username
	err := s.Configptr.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("user set to: %s\n", username)
	return nil
}
