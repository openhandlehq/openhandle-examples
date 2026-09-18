// Can I get X (Twitter) posts by date range?
// https://openhandle.dev/questions/can-i-get-x-posts-by-date-range
// Run: go run ./twitter/postssince

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	openhandle "github.com/openhandlehq/openhandle-go"
)

func main() {
	client, err := openhandle.New(os.Getenv("OPENHANDLE_TEST_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// Pagination stops once posts are older than Since. Filter the end date yourself.
	since := time.Date(2026, time.August, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2026, time.September, 1, 0, 0, 0, 0, time.UTC)

	posts := client.Twitter.Profile("copperfield_lab_x_test").Posts.Items(&openhandle.TwitterProfilePostsOptions{Since: &since})
	for posts.Next(ctx) {
		post := posts.Value()
		if post.CreatedAt.Before(until) {
			fmt.Println(post.CreatedAt, post.Text)
		}
	}
	if err := posts.Err(); err != nil {
		log.Fatal(err)
	}
}
