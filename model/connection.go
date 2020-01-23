package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	rdb *sql.DB
)

func ConnectDb(dsn string) {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
}

func ConnectReadonlyDb(dsn string) {
	var err error
	rdb, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
}

func CloseDb() {
	if db != nil {
		db.Close()
	}
}

func CloseReadonlyDb() {
	if rdb != nil {
		rdb.Close()
	}
}
