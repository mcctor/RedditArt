package reddit

import (
	"log"

	"github.com/turnage/graw/reddit"
)

var redditBot reddit.Bot

func init() {
	var err error
	redditBot, err = reddit.NewBot(reddit.BotConfig{
		Agent: "Ubuntu:github.com/mcctor/goreddit:v0.1.0(by /u/mcctor)",
		App: reddit.App{
			ID:       "K_P1LTKylZMqAw",
			Secret:   "snpKwH-hrKU29KeYFQxQn_wA9aQ",
			Username: "mcctor",
			Password: "@lienmwanga01",
		},
		Rate: 0,
	})
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
