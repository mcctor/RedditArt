package main

import (
	_ "github.com/mcctor/redditart/db"
	"github.com/mcctor/redditart/telegram"
)

func main() {
	telegram.Bot.Start()
}
