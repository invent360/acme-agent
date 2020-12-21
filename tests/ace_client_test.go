package tests

import (
	"context"
	"flag"
	"fmt"
	"github.com/acme-agent/swagger/client"
	"github.com/acme-agent/swagger/client/post"
	"log"
	"testing"
)

func TestSaveUser(t *testing.T) {

fmt.Println("Test triggered")
	flag.Parse()

	postClient := client.Default

	params := &post.ListParams{Context: context.Background()}

	posts, err := postClient.Post.List(params)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range posts.Payload {
		fmt.Printf("PostID: %v UserID: %v Title: %v Body: %v\n", p.ID, p.UserID, p.Title, p.Body)
		fmt.Println("-------------------------------------------------------------------------------")
	}
}
