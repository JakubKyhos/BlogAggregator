package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/JakubKyhos/blogaggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerRegister(s *state, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username required")
	}

	username := cmd.Args[0]

	var new_user = database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.db.CreateUser(context.Background(), new_user)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				fmt.Printf("user %s already exists\n", username)
				os.Exit(1)
			}
		}
		return err
	}

	s.configptr.CurrentUserName = username
	err = s.configptr.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("user: %s was created\n", username)
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
