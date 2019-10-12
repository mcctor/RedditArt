package reddit

import (
	"time"

	"github.com/mcctor/redditart/db"
	"github.com/mcctor/redditart/reddit/candidates"
	"github.com/mcctor/redditart/telegram"
	"github.com/turnage/graw/reddit"
)

const (
	historyLimit = 80
	timeToWait   = 30 // in minutes
	minutes      = 60
	minUpVotes   = 100
)

func isCandidate(upvotes int32) bool {
	if upvotes > minUpVotes {
		return true
	} else {
		return false
	}
}

// NewPosts periodically fetches new posts from given subreddit after stipulated delay
func NewPosts(subreddit string) {

	for {
		redditBot, err := reddit.NewBot(reddit.BotConfig{
			Agent: "Ubuntu:github.com/mcctor/redditart:v0.1.0(by /u/mcctor)",
			App: reddit.App{
				ID:       "K_P1LTKylZMqAw",
				Secret:   "snpKwH-hrKU29KeYFQxQn_wA9aQ",
				Username: "mcctor",
				Password: "@lienmwanga01",
			},
			Rate: 0,
		})
		if err != nil {
			continue
		}

		harvest, err := redditBot.Listing(subreddit, "")
		if err != nil {
			continue
		}

		candidatePosts := candidates.GetPosts()
		for _, post := range harvest.Posts[historyLimit:] {
			newPost := db.Post{
				PostID:   post.ID,
				Caption:  post.Title,
				Link:     post.Permalink,
				ImageUrl: post.URL,
				Author:   post.Author,
			}
			if _, exists := candidatePosts.Posts[newPost]; !exists && isCandidate(post.Ups) {
				candidatePosts.Posts[newPost] = struct{}{}
				db.AddPost(newPost)
				telegram.SendPhotoToAll(newPost)
			}
		}

		time.Sleep(timeToWait * minutes)
	}
}
