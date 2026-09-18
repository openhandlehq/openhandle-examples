// How do I get the view count of a TikTok video?
// https://openhandle.dev/questions/how-do-i-get-the-view-count-of-a-tiktok-video
// Run: go run ./tiktok/videoviews

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

	response, err := client.TikTok.Post("920100000000000001").Get(ctx, &openhandle.TikTokPostOptions{Freshness: openhandle.FreshnessLive})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(deref(response.Data.Metrics.Views), deref(response.Data.Metrics.Likes), response.CapturedAt) // 23000 1200 ...
}

func deref(value *int64) int64 {
	if value == nil {
		return -1 // null: the platform hid the number
	}
	return *value
}
