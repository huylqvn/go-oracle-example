package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func PostgresConnection(user, pass, host, port, dbName string) (*xorm.Engine, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbName)
	engine, err := xorm.NewEngine("postgres", dsn)
	if err != nil {
		fmt.Println("Error creating engine", err)
		return nil, err
	}
	fmt.Println("connection to Postgres successfully")
	return engine, nil
}
