package candidates

import (
	"github.com/mcctor/redditart/db"
)

// singleton object for candidates
var candidates *candidatePosts

type candidatePosts struct {
	Posts map[db.Post]struct{}
}

func GetPosts() *candidatePosts {
	if candidates != nil {
		return candidates
	} else {
		candidates = &candidatePosts{Posts: make(map[db.Post]struct{})}
		previousCandidates := db.GetAllPosts()

		// initialize singleton with database contents
		for _, previousCandidate := range previousCandidates {
			candidates.Posts[*previousCandidate] = struct{}{}
		}
		return candidates
	}
}
