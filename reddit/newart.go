package reddit

import (
	"github.com/mcctor/redditart/telegram"
	"time"

	"github.com/mcctor/redditart/db"
	"github.com/mcctor/redditart/reddit/candidates"
)

const (
	historyLimit  = 20
	waitTime      = 600
	minUpVotes    = 100
	postsToBeSent = 3
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
		candidatePosts := candidates.GetPosts()
		harvest, err := redditBot.Listing(subreddit, "")
		checkError(err)

		var count uint
		for _, post := range harvest.Posts[:historyLimit] {
			//if count < postsToBeSent {
			//	break
			//}
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
				count++
				telegram.SendPhotoToAll(newPost)
			}
		}

		time.Sleep(waitTime)
	}
}
