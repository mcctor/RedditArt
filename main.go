package main

import (
	_ "github.com/mcctor/redditart/db"
	"github.com/mcctor/redditart/reddit"
	"github.com/mcctor/redditart/telegram"
)

func main() {
	go reddit.NewPosts("/r/art")     // start as goroutine to fetch new posts
	go reddit.NewPosts("/r/artporn") // start as goroutine to fetch new posts
	telegram.Bot.Start()
}
