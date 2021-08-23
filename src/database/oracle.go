package database

import (
	"fmt"

	_ "github.com/mattn/go-oci8"

	"xorm.io/xorm"
)

func OracleConnection(dsn string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("oci8", dsn)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("connection to Oracle successfully")

	return engine, nil
}
