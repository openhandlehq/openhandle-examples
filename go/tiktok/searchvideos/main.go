// How do I search TikTok videos by keyword?
// https://openhandle.dev/questions/how-do-i-search-tiktok-videos-by-keyword
// Run: go run ./tiktok/searchvideos

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

	videos := client.TikTok.Search.Posts.Items(&openhandle.TikTokSearchPostsOptions{Q: "synthetic"})
	for videos.Next(ctx) {
		video := videos.Value()
		fmt.Println(video.Author.Handle, deref(video.Metrics.Views))
	}
	if err := videos.Err(); err != nil {
		log.Fatal(err)
	}
}

func deref(value *int64) int64 {
	if value == nil {
		return -1 // null: the platform hid the number
	}
	return *value
}
