// How do I get a Reddit user's post history?
// https://openhandle.dev/questions/how-do-i-get-a-reddit-users-post-history
// Run: go run ./reddit/userposthistory

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openhandle "github.com/openhandlehq/openhandle-go"
)

func main() {
	client, err := openhandle.New(os.Getenv("OPENHANDLE_TEST_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	posts := client.Reddit.Profile("synthetic_reddit").Posts.Items(&openhandle.RedditProfilePostsOptions{Sort: openhandle.SortOrder("new")})
	for posts.Next(ctx) {
		post := posts.Value()
		fmt.Println(post.Community.DisplayHandle, post.CreatedAt, post.Title)
	}
	if err := posts.Err(); err != nil {
		log.Fatal(err)
	}
}
