package repository

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db          *sql.DB
	errNotFound = errors.New("account not found")
)

// Open opens a database
func Open() {
	sqlite3Database, err := sql.Open("sqlite3", "./sqlite/bank.db")
	db = sqlite3Database
	exitErr(err)
	log.Println("sqlite3 database opened")
}

// Close closes the database
func Close() {
	log.Println("sqlite3 database closed")
	db.Close()
}

func exitErr(err error) {
	if err != nil {
		Close()
		panic(err)
	}
}
