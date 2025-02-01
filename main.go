package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/JakubKyhos/blogaggregator/internal/config"
	"github.com/JakubKyhos/blogaggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db        *database.Queries
	configptr *config.Config
}

func main() {
	dbURL := "postgres://postgres:postgres@localhost:5432/gator"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	var s = &state{
		db:        dbQueries,
		configptr: &cfg,
	}

	commands := Commands{
		Handlers: make(map[string]func(*state, Command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsersList)
	commands.register("agg", handlerAggCmd)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerListFeeds)
	commands.register("follow", middlewareLoggedIn(handlerfollow))
	commands.register("following", middlewareLoggedIn(handlerfollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerunfollow))

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Fprintf(os.Stderr, "Error: command name required\n")
		os.Exit(1)
	}

	var command = Command{
		Name: arguments[1],
		Args: arguments[2:],
	}

	err = commands.Run(s, command)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
