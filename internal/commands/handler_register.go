package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/levabu/gator/internal/database"
	"github.com/levabu/gator/internal/state"
)

func HandlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}
	username := cmd.Args[0]

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: username,
	})

	if err != nil {
		if strings.Contains(fmt.Sprint(err), "pq: duplicate key value violates unique constraint") {
			return fmt.Errorf("user %s already exists", username)
		}
		return fmt.Errorf("error when creating a user: %w", err)
	}

	fmt.Printf("User was created:\n%+v", user)

	s.Config.SetUser(username)
	return nil
}
