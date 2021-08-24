package database

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

func SqliteConnection() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", "data/backup.db")
	if err != nil {
		fmt.Println("Error creating engine", err)
		return nil, err
	}
	fmt.Println("connection to Sqlite3 successfully")
	return engine, nil
}
