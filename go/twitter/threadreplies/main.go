// How do I get replies or a full thread on X (Twitter)?
// https://openhandle.dev/questions/how-do-i-get-replies-or-a-full-thread-on-x
// Run: go run ./twitter/threadreplies

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

	replies := client.Twitter.Post("940100000000000001").Comments.Items(nil)
	for replies.Next(ctx) {
		reply := replies.Value()
		fmt.Println(reply.Author.Handle, reply.Text)
	}
	if err := replies.Err(); err != nil {
		log.Fatal(err)
	}
}
