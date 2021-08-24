package database

import (
	"fmt"

	_ "github.com/mattn/go-oci8"

	"xorm.io/xorm"
)

func OracleConnection(user, pass, host, port, dbName string) (*xorm.Engine, error) {
	dsn := fmt.Sprintf("%s/%s@%s:%s/%s", user, pass, host, port, dbName)
	engine, err := xorm.NewEngine("oci8", dsn)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("connection to Oracle successfully")

	return engine, nil
}
