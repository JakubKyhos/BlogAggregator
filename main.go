package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JakubKyhos/blogaggregator/internal/config"
)

type state struct {
	Configptr *config.Config
}

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	var s = &state{
		Configptr: &cfg,
	}

	commands := Commands{
		Handlers: make(map[string]func(*state, Command) error),
	}

	commands.Register("login", HandlerLogin)

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
