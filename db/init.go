package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	dbDialect = "sqlite3"
	dbFile    = "db.sqlite3"
)

func init() {
	db := createSess(dbDialect, dbFile)
	defer db.Close()

	// migrate models to database
	db.AutoMigrate(&Post{}, &User{})
}

func createSess(dbDialect string, dbFile string) *gorm.DB {
	db, err := gorm.Open(dbDialect, dbFile)
	checkError(err)
	return db
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
