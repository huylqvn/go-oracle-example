package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("data/backup.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
