package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func MysqlConnection(user, pass, host, port, dbName string) (*xorm.Engine, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		dbName)

	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("Error creating engine", err)
		return nil, err
	}
	fmt.Println("connection to Mysql successfully")
	return engine, nil
}
