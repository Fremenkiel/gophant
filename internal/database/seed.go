package database

import (
	"database/sql"
	"log"

	"github.com/Fremenkiel/gophant/v2/internal/models"
)

func SeedData() {
	db := CurrentDB()

	seedGroups(db)
}

func seedGroups(db *sql.DB) {
	groups := []models.Group{
		{ ID: 1, Name: "Production", R: 226, G: 107, B: 99 },
		{ ID: 2, Name: "Staging", R: 229, G: 180, B: 100 },
		{ ID: 3, Name: "Dev", R: 115, G: 193, B: 136 },
	}

	for _, obj := range groups {
		var res sql.NullInt16
		err := db.QueryRow("SELECT COUNT(*) FROM groups WHERE name = ?;", obj.Name).Scan(&res)
		if err != nil {
			log.Fatal(err)
		}

		if res.Valid && res.Int16 == 0 {
			stmt, err := db.Prepare("INSERT INTO groups (ID, Name, R, G, B) VALUES (?, ?, ?, ?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(obj.ID, obj.Name, obj.R, obj.G, obj.B)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
