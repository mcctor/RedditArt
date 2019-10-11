package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mcctor/redditart/db"
	"github.com/mcctor/redditart/reddit"
	"github.com/mcctor/redditart/telegram"
)

func main() {
	if !(len(os.Args) > 1) {
		log.Fatal("Need to provide subreddit names")
	}
	for _, subreddit := range os.Args[1:] {
		go func() {
			fmtSubredditName := fmt.Sprintf("/r/%s/", strings.ToLower(subreddit))
			reddit.NewPosts(fmtSubredditName)
		}()
	}
	telegram.Bot.Start()
}
