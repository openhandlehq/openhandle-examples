// How do I get Instagram hashtag posts past the 30 hashtag limit?
// https://openhandle.dev/questions/how-do-i-get-instagram-hashtag-posts-past-the-30-hashtag-limit
// Run: go run ./instagram/hashtagposts

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

	posts := client.Instagram.Hashtag("syntheticvolume").Posts.Items(&openhandle.InstagramHashtagPostsOptions{Sort: openhandle.SortOrder("recent")})
	for posts.Next(ctx) {
		post := posts.Value()
		fmt.Println(post.ID, post.Author.Handle, deref(post.Metrics.Likes))
	}
	if err := posts.Err(); err != nil {
		log.Fatal(err)
	}
}

func deref(value *int64) int64 {
	if value == nil {
		return -1 // null: the platform hid the number
	}
	return *value
}
