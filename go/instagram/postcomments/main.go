// How do I get comments on an Instagram post that is not mine?
// https://openhandle.dev/questions/how-do-i-get-comments-on-an-instagram-post-that-is-not-mine
// Run: go run ./instagram/postcomments

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

	comments := client.Instagram.Post("910100000001").Comments.Items(nil)
	for comments.Next(ctx) {
		comment := comments.Value()
		fmt.Println(comment.Author.Handle, comment.Text)
	}
	if err := comments.Err(); err != nil {
		log.Fatal(err)
	}
}
