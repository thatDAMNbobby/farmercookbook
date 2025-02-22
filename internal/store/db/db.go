package db

import (
	"farmercookbook/internal/store"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func open(dbName string) (*gorm.DB, error) {

	// make the temp directory if it doesn't exist
	err := os.MkdirAll("/tmp", 0755)
	if err != nil {
		return nil, err
	}

	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}

func MustOpen(dbName string) *gorm.DB {

	if dbName == "" {
		dbName = "cookbook.db"
	}

	db, err := open(dbName)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&store.User{}, &store.Session{}, &store.Recipe{})
	if err != nil {
		panic(err)
	}

	return db
}
