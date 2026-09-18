// How do I track brand mentions across subreddits?
// https://openhandle.dev/questions/how-do-i-track-brand-mentions-across-subreddits
// Run: go run ./reddit/brandmentions

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

	// Run on a schedule. Store post IDs and skip the ones you have seen.
	day := "day"
	page, err := client.Reddit.Search.Posts.List(ctx, &openhandle.RedditSearchPostsOptions{Q: "synthetic", Sort: openhandle.SortOrder("new"), T: &day})
	if err != nil {
		log.Fatal(err)
	}

	for _, post := range page.Data {
		fmt.Println(post.Community.DisplayHandle, deref(post.Metrics.Score), post.Title)
	}
}

func deref(value *int64) int64 {
	if value == nil {
		return -1 // null: the platform hid the number
	}
	return *value
}
