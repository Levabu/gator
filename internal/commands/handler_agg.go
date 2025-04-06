package commands

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/levabu/gator/internal/database"
	"github.com/levabu/gator/internal/rss"
	"github.com/levabu/gator/internal/state"
)

func HandlerAgg(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}

	time_between_reqs := cmd.Args[0]
	duration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("invalid time string format: %w", err)
	}

	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for ; ; <- ticker.C {
		go scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state.State) error {
	nextFeed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}
	err = s.DB.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	fmt.Printf("Fetch request for %s\n", nextFeed.Url)
	fetchedFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}
	savePosts(s, fetchedFeed, nextFeed.ID)
	return nil
}

func savePosts(s *state.State, feed *rss.RSSFeed, feedId uuid.UUID) error {
	for _, item := range feed.Channel.Item {
		date := sql.NullTime{}
		if v, err := time.Parse(time.RFC1123, item.PubDate); err == nil {
			date = sql.NullTime{
				Time: v,
				Valid: err == nil,
			}
		}
		_, err := s.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Description: item.Description,
			Url: item.Link,
			PublishedAt: date,
			FeedID: feedId,
		})
		if err != nil {
			if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
				continue
			}
			fmt.Printf("error creating post: %s\n", err)
		}
	}
	return nil
}