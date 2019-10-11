package reddit

import (
	"github.com/mcctor/redditart/db"
	"github.com/mcctor/redditart/reddit/candidates"
	"github.com/mcctor/redditart/telegram"
)

const (
	historyLimit = 80
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
	candidatePosts := candidates.GetPosts()
	harvest, err := redditBot.Listing(subreddit, "")
	checkError(err)

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
}
