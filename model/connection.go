package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func ConnectDb(dsn string) {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
}

func CloseDbConnection() {
	if db != nil {
		db.Close()
	}
}
