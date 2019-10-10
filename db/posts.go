package db

type Post struct {
	PostID   string `gorm:"unique"`
	Caption  string
	Link     string
	ImageUrl string
	Author   string
}

func AddPost(newPost Post) {
	db := createSess(dbDialect, dbFile)
	defer db.Close()

	db.Create(&newPost)
}

func GetPostByID(id int) (*Post, bool) {
	db := createSess(dbDialect, dbFile)
	defer db.Close()

	var foundPost Post
	db.First(&foundPost, "post_id = ?", id)
	if foundPost.PostID == "" {
		return nil, false
	} else {
		return &foundPost, true
	}
}

func GetAllPosts() (allPosts []*Post) {
	db := createSess(dbDialect, dbFile)
	defer db.Close()

	db.Find(&allPosts)
	return
}
