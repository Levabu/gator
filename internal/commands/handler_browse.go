package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/levabu/gator/internal/database"
	"github.com/levabu/gator/internal/state"
)

func HandlerBrowse(s *state.State, cmd Command, user database.User) error {
	var postsLimit int32 = 2
	if len(cmd.Args) > 0 {
		parsedLimit, parseErr := strconv.ParseInt(cmd.Args[0], 10, 32)
		if parseErr != nil {
			return fmt.Errorf("usage: <posts_number>")
		}
		postsLimit = int32(parsedLimit)
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: postsLimit,
	})
	fmt.Printf("Number of posts: %v\n", len(posts))
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)		
	}
	for i, post := range posts {
		fmt.Printf("%v) %s\n", i, post.Title)
		fmt.Printf("%s\n", post.PublishedAt.Time)
		fmt.Printf("%s\n", post.Url)
		fmt.Printf("%s\n\n", post.Description)
	}
	
	return nil
}
