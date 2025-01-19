package main

import (
	"context"
	"fmt"

	"github.com/chaseplamoureux/blogaggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		currentUsername := s.ConfPointer.Username
		currentUser, err := s.dbConn.GetUser(context.Background(), currentUsername)
		if err != nil {
			return fmt.Errorf("User not found")
		}
		return handler(s, c, currentUser)
	}
		

}

