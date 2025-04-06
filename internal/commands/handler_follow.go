package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/levabu/gator/internal/database"
	"github.com/levabu/gator/internal/state"
)

func HandlerFollow(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed by url: %w", err)
	}

	feedFollow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Printf("User %s now follows %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}