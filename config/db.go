package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var BaseURL = "http://localhost:8080"

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/todolist?Pard")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(60 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)

	return db
}
