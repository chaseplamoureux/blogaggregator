package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/chaseplamoureux/blogaggregator/internal/database"
	"github.com/google/uuid"
)

type command struct {
	commandName string
	commandArgs []string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, exists := c.commandsMap[cmd.commandName]
	if !exists {
		return fmt.Errorf("Command not found %s", cmd.commandName)
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil

}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.commandArgs) == 0 {
		return errors.New("No username was provided")
	}

	username := cmd.commandArgs[0]
	_, err := s.dbConn.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("User does not exist: %v\n", err)
		os.Exit(1)
	}

	s.ConfPointer.SetUser(username)

	fmt.Println("username has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.commandArgs) == 0 {
		return errors.New("No username was provided")
	}
	username := cmd.commandArgs[0]
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	}
	registeredUser, err := s.dbConn.CreateUser(context.Background(), newUser)
	if err != nil {
		fmt.Printf("error occurred creating new user in DB: %v\n", err)
		os.Exit(1)
	}
	s.ConfPointer.SetUser(registeredUser.Name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.dbConn.DeleteUsers(context.Background())
	if err != nil {
		fmt.Printf("Error occurred deleting users from table\n")
		return err
	}

	fmt.Println("Users table has been reset")
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.dbConn.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user == s.ConfPointer.Username {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.commandArgs) == 0 {
	// 	return errors.New("No feed URL was provided")
	// }
	// url := cmd.commandArgs[0]
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	fmt.Printf("%v\n", rssFeed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.commandArgs) != 2 {
		return fmt.Errorf("Invalid number of required arguments")
	}

	currentUsername := s.ConfPointer.Username
	currentuser, err := s.dbConn.GetUser(context.Background(), currentUsername)
	if err != nil {
		return fmt.Errorf("User not found")
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      cmd.commandArgs[0],
		Url:       cmd.commandArgs[1],
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentuser.ID,
	}

	feed, err := s.dbConn.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return fmt.Errorf("Error writing feed to DB: %v", err)
	}

	fmt.Println("Feed added to DB")
	printFeed(feed)
	return nil

}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:              %s\n", feed.ID)
	fmt.Printf("* Name:            %s\n", feed.Name)
	fmt.Printf("* URL:             %s\n", feed.Url)
	fmt.Printf("* Created:         %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:         %v\n", feed.UpdatedAt)
	fmt.Printf("* UserID:          %s\n", feed.UserID)
}

func handlerFeeds(s *state, cmd command) error {

	feeds, err := s.dbConn.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.dbConn.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println("----------")
		fmt.Printf("%s\n", feed.Name)
		fmt.Printf("%s\n", feed.Url)
		fmt.Printf("%s\n", user.Name)
		fmt.Println("----------")
		fmt.Println("")
	}
	return nil
}
