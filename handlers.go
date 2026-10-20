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
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("User %s does not exist.", username)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Error while logging in: %w\n", err)
	}

	fmt.Printf("User %s logged in.\n", user.Name)
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

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Error while logging in: %w\n", err)
	}

	fmt.Printf("User %s logged in.\n", user.Name)

	return nil
}

func handleReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("Could not delete data from users table: %w\n", err)
	}

	fmt.Printf("users table data deleted successfully.\n")

	return nil
}

func handleUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error while getting users: %w\n", err)
	}

	currentUser := s.config.CurrentUserName

	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}

	return nil
}

func handleAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	rss, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Print(rss)

	return nil
}

func handleAddFeed(s *state, cmd command) error {
	if cmd.args == nil {
		return fmt.Errorf("Please provide a feed to add.\n")
	}

	if len(cmd.args) != 2 {
		return fmt.Errorf("Please provide a feed name and URL.\n")
	}

	name, url := cmd.args[0], cmd.args[1]

	cxt := context.Background()
	currentUserName := s.config.CurrentUserName
	currentUser, err := s.db.GetUser(cxt, currentUserName)
	if err != nil {
		return fmt.Errorf("Error while getting current user data from the db.\n")
	}

	now := sql.NullTime{Time: time.Now(), Valid: true}
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	}
	feed, err := s.db.CreateFeed(cxt, feedParams)
	if err != nil {
		return fmt.Errorf("Error while creating feed: %w\n", err)
	}

	fmt.Printf("Feed %s created successfully for user %s.\n", feed.Name, currentUserName)

	return nil
}
