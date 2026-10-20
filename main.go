package main

import (
	"blogaggr/internal"
	"blogaggr/internal/database"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *internal.Config
}

func main() {
	config, err := internal.Read()
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", config.DbURL)
	dbQueries := database.New(db)

	state := state{config: &config, db: dbQueries}

	commands := NewCommands()
	commands.register("login", handleLogin)
	commands.register("register", handleRegister)
	commands.register("reset", handleReset)
	commands.register("users", handleUsers)
	commands.register("agg", handleAgg)
	commands.register("addfeed", handleAddFeed)

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
