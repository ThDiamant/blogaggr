package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmdRegistry map[string]func(s *state, cmd command) error
}

func NewCommands() *commands {
	var cmds commands
	cmds.cmdRegistry = make(map[string]func(s *state, cmd command) error)

	return &cmds
}

func (c *commands) run(s *state, cmd command) error {
	callback, ok := c.cmdRegistry[cmd.name]
	if !ok {
		return fmt.Errorf("The given command does not exist.\n")
	}

	err := callback(s, cmd)
	if err != nil {
		return fmt.Errorf("Error executing %s command: %w\n", cmd.name, err)
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdRegistry[name] = f
}
