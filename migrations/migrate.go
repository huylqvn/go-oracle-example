package main

import (
	"database/sql"
	"flag"
	"fmt"
	"go-oracle/config"

	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateUp(driver, dbUrl string) bool {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/db",
	}
	db, err := sql.Open(driver, dbUrl)
	if err != nil {
		fmt.Println("Error Connection", err)
		return false
	}

	n, err := migrate.Exec(db, driver, migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error migration", err)
		return false
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return true
}

func MigrateDown(driver, dbUrl string) bool {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/db",
	}
	db, err := sql.Open(driver, dbUrl)
	if err != nil {
		fmt.Println("Error Connection", err)
		return false
	}

	n, err := migrate.Exec(db, driver, migrations, migrate.Down)
	if err != nil {
		fmt.Println("Error migration", err)
		return false
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return true
}

func main() {
	godotenv.Load()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBName,
		cfg.DBPass)

	var t string
	flag.StringVar(&t, "t", "", "choose seed type\nt=up || t=down")
	flag.Parse()

	if t == "up" {
		if !MigrateUp("mysql", dbUrl) {
			fmt.Println("Start service error with migrations")
		}
	} else if t == "down" {
		if !MigrateDown("mysql", dbUrl) {
			fmt.Println("Start service error with migrations")
		}
	}
}
