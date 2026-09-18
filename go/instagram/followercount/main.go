// How do I get an Instagram follower count without logging in?
// https://openhandle.dev/questions/how-do-i-get-an-instagram-follower-count-without-logging-in
// Run: go run ./instagram/followercount

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

	response, err := client.Instagram.Profile("northstar_forge_test").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Data.Handle, deref(response.Data.Metrics.Followers), response.CapturedAt) // northstar_forge_test 48291 ...
}

func deref(value *int64) int64 {
	if value == nil {
		return -1 // null: the platform hid the number
	}
	return *value
}
