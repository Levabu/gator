package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/levabu/gator/internal/database"
	"github.com/levabu/gator/internal/state"
)

func HandlerAddFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %v <feed> <url>", cmd.Name)
	}
	feedName, feedUrl := cmd.Args[0], cmd.Args[1]

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: feedName,
		Url: feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: feed.UserID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Printf("Feed created:\n")
	fmt.Printf("ID: %v\n", feed.ID)
	fmt.Printf("Name: %v\n", feed.Name)
	fmt.Printf("URL: %v\n", feed.Url)

	return nil
}
