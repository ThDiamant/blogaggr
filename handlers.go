package main

import (
	"blogaggr/internal/database"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handleLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return fmt.Errorf("Please provide a user to login.\n")
	}

	username := cmd.args[0]
	err := s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("Error while logging in: %w\n", err)
	}

	fmt.Printf("User %s logged in.\n", username)
	return nil
}

func handleRegister(s *state, cmd command) error {
	if cmd.args == nil {
		return fmt.Errorf("Please provide a user to register.\n")
	}

	now := sql.NullTime{Time: time.Now(), Valid: true}
	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.args[0],
	}
	user, err := s.db.CreateUser(context.Background(), createUserParams)
	if err != nil {
		return fmt.Errorf("Error while registering user: %w\n", err)
	}

	fmt.Printf("User %s registered successfully.\n", user.Name)
	fmt.Println(user)

	return nil
}
