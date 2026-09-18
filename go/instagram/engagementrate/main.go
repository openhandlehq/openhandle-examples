// How do I get a competitor's Instagram engagement rate?
// https://openhandle.dev/questions/how-do-i-get-a-competitors-instagram-engagement-rate
// Run: go run ./instagram/engagementrate

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

	profile := client.Instagram.Profile("northstar_forge_test")
	account, err := profile.Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	page, err := profile.Posts.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Skip posts with hidden likes. Nil is not zero.
	var total, count int64
	for _, post := range page.Data {
		if post.Metrics.Likes == nil || post.Metrics.Comments == nil {
			continue
		}
		total += *post.Metrics.Likes + *post.Metrics.Comments
		count++
	}
	rate := float64(total) / float64(count) / float64(*account.Data.Metrics.Followers) * 100

	fmt.Printf("%.2f%% over %d posts\n", rate, count)
}
