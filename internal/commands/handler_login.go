package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/levabu/gator/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	username := cmd.Args[0]

	user, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "no rows in result set") {
			return fmt.Errorf("user with name %s doesn't exist", username)
		}
		return fmt.Errorf("error getting a user %s: %w", username, err)
	}

	if err := s.Config.SetUser(user.Name); err != nil {
		return fmt.Errorf("error logging in: %w", err)
	}
	
	fmt.Printf("Logged in as %s\n", user.Name)

	return nil
}
