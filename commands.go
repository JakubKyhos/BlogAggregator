package main

import (
	"fmt"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*state, Command) error
}

func (c *Commands) register(name string, f func(*state, Command) error) {
	c.Handlers[name] = f
}

func (c *Commands) Run(s *state, cmd Command) error {
	function, exists := c.Handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("command %s does not exists", cmd.Name)
	}
	return function(s, cmd)
}
