package persistence

import (
	"database/sql"
	"errors"
	"log"
)

var (
	db          *sql.DB
	errNotFound = errors.New("account not found")
)

func OpenSqlite3() {
	sqlite3Database, err := sql.Open("sqlite3", "./infrastructure/sqlite/bank.db")
	if err != nil {
		panic(err)
	}
	db = sqlite3Database
	log.Println("sqlite3 database opened")
}

func CloseSqlite3() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
	log.Println("sqlite3 database closed")
}
