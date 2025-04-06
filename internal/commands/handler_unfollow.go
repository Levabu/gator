package commands

import (
	"context"
	"fmt"

	"github.com/levabu/gator/internal/database"
	"github.com/levabu/gator/internal/state"
)

func HandlerUnfollow(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed by url: %w", err)
	}

	err = s.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error deleting feed follow: %w", err)
	}

	fmt.Printf("User %s no longer follows %s\n", feed.Name, user.Name)

	return nil
}