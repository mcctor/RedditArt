package reddit

import (
	"log"

	"github.com/turnage/graw/reddit"
)

var redditBot reddit.Bot

func init() {
	var err error
	redditBot, err = reddit.NewBotFromAgentFile(".agent", 0)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
