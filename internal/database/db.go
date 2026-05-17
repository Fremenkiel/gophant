package database

import (
	"database/sql"
	"log"
	"os"
	"sync/atomic"
)

var perDb atomic.Pointer[sql.DB]

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DB_NAME")+"?_journal=WAL&_fk=on")
	if err != nil { return nil, err }
	if err := db.Ping(); err != nil { return nil, err }
	perDb.Store(db)
	return db, nil
}

func CurrentDB() *sql.DB {
	db := perDb.Load()
	if db == nil {
		log.Fatal("DB not loaded.")
	}
	return db
}
