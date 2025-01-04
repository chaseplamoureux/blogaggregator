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
	s.ConfPointer.SetUser(username)

	fmt.Println("username has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.commandArgs) == 0 {
		return errors.New("No username was provided")
	}
	newUser := database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.commandArgs[0]}
	_, err := s.dbConn.CreateUser(context.Background(), newUser)
	if err != nil {
		fmt.Printf("error occurred creating new user in DB: %v\n", err)
		os.Exit(1)
	}
	return nil
}
