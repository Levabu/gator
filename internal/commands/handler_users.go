package commands

import (
	"context"
	"fmt"

	"github.com/levabu/gator/internal/state"
)

func HandlerUsers(s *state.State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}

	currentUserName := s.Config.CurrentUserName
	for _, user := range users {
		isCurrent := user.Name == currentUserName
		f := "* %s"
		if isCurrent {
			f += " (current)"
		}
		fmt.Printf(f + "\n", user.Name)
	}

	return nil
}
