package database

import (
	"database/sql"
	"log"
	"os"
)

func CreateDB() *sql.DB {
	db, err := sql.Open("sqlite3", os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
