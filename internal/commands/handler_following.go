package commands

import (
	"context"
	"fmt"

	"github.com/levabu/gator/internal/database"
	"github.com/levabu/gator/internal/state"
)

func HandlerFollowing(s *state.State, cmd Command, user database.User) error {
	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feeds by user_id: %w", err)
	}

	if len(feeds) == 0 {
		return nil
	}

	fmt.Printf("User %s follows:\n", user.Name)
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}