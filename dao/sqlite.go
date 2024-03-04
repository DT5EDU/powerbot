package dao

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDatabase(database string) *DB {
	db, err := gorm.Open(sqlite.Open((database)), &gorm.Config{})
	if err != nil {
		panic("db init failed")
	}

	db.AutoMigrate(&Group{})

	return &DB{db: db}
}
