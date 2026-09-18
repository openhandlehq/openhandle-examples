// How do I search X (Twitter) posts by keyword?
// https://openhandle.dev/questions/how-do-i-search-x-twitter-posts-by-keyword
// Run: go run ./twitter/searchposts

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

	posts := client.Twitter.Search.Posts.Items(&openhandle.TwitterSearchPostsOptions{Q: "synthetic", Sort: openhandle.SortOrder("latest")})
	for posts.Next(ctx) {
		post := posts.Value()
		fmt.Println(post.Author.Handle, post.Text)
	}
	if err := posts.Err(); err != nil {
		log.Fatal(err)
	}
}
