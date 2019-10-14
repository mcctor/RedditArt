package reddit

import (
	"log"
	"os"
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
			Agent: os.Getenv("REDDIT_AGENT"),
			App: reddit.App{
				ID:       os.Getenv("REDDIT_ID"),
				Secret:   os.Getenv("REDDIT_SECRET"),
				Username: os.Getenv("REDDIT_USERNAME"),
				Password: os.Getenv("REDDIT_PASSWORD"),
			},
			Rate: 0,
		})
		if err != nil {
			log.Fatal(err)
		}

		harvest, err := redditBot.Listing(subreddit, "")
		if err != nil {
			log.Fatal(err)
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
