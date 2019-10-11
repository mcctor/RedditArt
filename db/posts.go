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

func GetAllPosts() (allPosts []*Post) {
	db := createSess(dbDialect, dbFile)
	defer db.Close()

	db.Find(&allPosts)
	return
}
