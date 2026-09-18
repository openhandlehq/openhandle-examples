// How do I get a TikTok follower count for any username?
// https://openhandle.dev/questions/how-do-i-get-a-tiktok-follower-count-for-any-username
// Run: go run ./tiktok/followercount

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

	response, err := client.TikTok.Profile("pixel_orchard_tt_test").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Data.Handle, deref(response.Data.Metrics.Followers), deref(response.Data.Metrics.Likes)) // pixel_orchard_tt_test 75120 9814220
}

func deref(value *int64) int64 {
	if value == nil {
		return -1 // null: the platform hid the number
	}
	return *value
}
