package main

import (
	"errors"
	"fmt"
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
	f(s, cmd)
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