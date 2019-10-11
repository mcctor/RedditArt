package db

type User struct {
	UserID    int `gorm:"unique"`
	FirstName string
	Streaming bool `gorm:"default:false"`
}

func UpdateUser(user *User) {
	db := createSess(dbDialect, dbFile)
	defer db.Close()

	db.Model(user).Updates(*user)
}

func AddUser(user User) {
	db := createSess(dbDialect, dbFile)
	defer db.Close()

	db.Create(&user)
}

func GetUserByID(id int) (*User, bool) {
	db := createSess(dbDialect, dbFile)
	defer db.Close()

	var foundUser User
	db.First(&foundUser, "user_id = ?", id)
	if foundUser.FirstName != "" {
		return &foundUser, true
	} else {
		return nil, false
	}
}

func GetAllUsers(streaming bool) (allUsers []*User) {
	db := createSess(dbDialect, dbFile)
	defer db.Close()

	db.Where("streaming = ?", streaming).Find(&allUsers)
	return
}
