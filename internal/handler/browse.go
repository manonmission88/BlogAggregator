package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// return all the feeds in the database
func HandlerBrowse(s *state.State, cmd Command, user database.User) error {

	// if no arg is provided , set the default limit as 2
	var val int32 = 2 // Default limit
	if len(cmd.Args) > 0 {
		if intVal, err := strconv.Atoi(cmd.Args[0]); err == nil && intVal > 0 {
			val = int32(intVal)
		} else if err != nil {
			return fmt.Errorf("invalid limit argument: %v", err)
		} else {
			return fmt.Errorf("limit must be a positive integer")
		}
	}

	posts, err := s.DbQueries.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  val,
	})
	if err != nil {
		return fmt.Errorf("could not get the post %w\n", err)
	}
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("post titles: %s \n", post.Title)
		fmt.Printf("post published at: %s \n", post.PublishedAt.Time)
		fmt.Printf("post created at: %s \n", post.CreatedAt)
		fmt.Printf("post description: %s \n", post.Description.String)
		fmt.Printf("post url: %s \n", post.Url)
		println("===============================")

	}
	return nil
}
