package main

import (
	"github.com/mcctor/redditart/reddit"
	"github.com/mcctor/redditart/telegram"
)

func main() {
	go reddit.NewPosts("/r/art")
	go reddit.NewPosts("/r/artporn")

	telegram.Bot.Start()
}
