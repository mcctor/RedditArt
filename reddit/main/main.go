package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/mcctor/redditart/reddit"
)

var wg sync.WaitGroup

func main() {
	if !(len(os.Args) > 1) {
		log.Fatal("Need to provide subreddit names")
	}
	for _, subreddit := range os.Args[1:] {
		wg.Add(1)
		go func() {
			fmtSubredditName := fmt.Sprintf("/r/%s/", strings.ToLower(subreddit))
			reddit.NewPosts(fmtSubredditName)
			wg.Done()
		}()
	}

	wg.Wait()
}
