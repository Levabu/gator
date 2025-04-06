package commands

import (
	"context"
	"fmt"

	"github.com/levabu/gator/internal/state"
)

func HandlerReset(s *state.State, cmd Command) error {
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting users: %w", err)
	}
	
	fmt.Println("Database reset: all users deletetd")

	return nil
}
