package main

import (
	"context"
	"fmt"

	"github.com/levabu/gator/internal/commands"
	"github.com/levabu/gator/internal/database"
	"github.com/levabu/gator/internal/state"
)


func middlewareLoggedIn(handler func(s *state.State, cmd commands.Command, user database.User) error) func(*state.State, commands.Command) error {
	return func(s *state.State, c commands.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting a user %s: %w", s.Config.CurrentUserName, err)
		}
		err = handler(s, c, user)
		if err != nil {
			return err
		}
		return nil
	}
}