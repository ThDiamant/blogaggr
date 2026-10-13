package main

import (
	"blogaggr/internal"
	"fmt"
	"os"
)

type state struct {
	config *internal.Config
}

type command struct {
	name string
	args []string
}

func handleLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return fmt.Errorf("Please provide a user to login.")
	}

	username := cmd.args[0]
	err := s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("Error while logging in: %w\n", err)
	}

	fmt.Printf("User %s logged in.\n", username)
	return nil
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

func main() {
	config, err := internal.Read()
	if err != nil {
		fmt.Println(err)
	}

	state := state{config: &config}

	commands := NewCommands()
	commands.register("login", handleLogin)

	inputArgs := os.Args
	if len(inputArgs) < 2 {
		fmt.Println("Please provide more than 1 argument.")
		os.Exit(1)
	}

	var cmdArgs []string
	if len(inputArgs) > 2 {
		cmdArgs = inputArgs[2:]
	}

	cmd := command{
		name: inputArgs[1],
		args: cmdArgs,
	}

	err = commands.run(&state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config, err = internal.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(config)
}
