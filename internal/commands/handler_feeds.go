package commands

import (
	"context"
	"fmt"

	"github.com/levabu/gator/internal/state"
)

func HandlerFeeds(s *state.State, cmd Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error gettings feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("Created by: %s\n\n", feed.UserName)
	}
	return nil
}