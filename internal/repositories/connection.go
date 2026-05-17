package repositories

import (
	"errors"

	"github.com/Fremenkiel/gophant/v2/internal/database"
	"github.com/Fremenkiel/gophant/v2/internal/models"
)

type ConnectionRepository struct {
}

func NewConnectionRepository() *ConnectionRepository {
	return &ConnectionRepository{}
}

func (r *ConnectionRepository) GetAll() ([]models.Connection, error) {
	db := database.CreateDB()

	rows, err := db.Query("SELECT * FROM connections")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []models.Connection

	for rows.Next() {
		var c models.Connection
		if err := rows.Scan(&c.ID, &c.Name, &c.Permission, &c.Address, &c.Port, &c.Database, &c.Username, &c.Password, &c.Status); err != nil {
			return nil, err
		}
		connections = append(connections, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return connections, nil
}

func (r *ConnectionRepository) Get(id uint) (models.Connection, error) {
	db := database.CreateDB()

	rows, err := db.Query("SELECT * FROM connections WHERE id = ?")
	if err != nil {
		return models.Connection{}, err
	}
	defer rows.Close()

	var c models.Connection

	for rows.Next() {
		if err := rows.Scan(&c.ID, &c.Name, &c.Permission, &c.Address, &c.Port, &c.Database, &c.Username, &c.Password, &c.Status); err != nil {
			return models.Connection{}, err
		}
	}
	if err = rows.Err(); err != nil {
		return models.Connection{}, err
	}

	return c, nil
}

func (r *ConnectionRepository) Create(connection models.Connection) error {
	db := database.CreateDB()

	stmt, err := db.Prepare("INSERT INTO connections (name, permission, address, port, database, username, password, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(connection.Name, connection.Permission, connection.Address, connection.Port, connection.Database, connection.Username, connection.Password, connection.Status)
	return err
}

func (r *ConnectionRepository) Update(connection models.Connection) error {
	db := database.CreateDB()

	stmt, err := db.Prepare("UPDATE connections SET name = ?, permission = ?, address = ?, port = ?, database = ?, username = ?, password = ?, status = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(connection.Name, connection.Permission, connection.Address, connection.Port, connection.Database, connection.Username, connection.Password, connection.Status, connection.ID)
	return err
}

func (r *ConnectionRepository) Delete(connection models.Connection) error {
	db := database.CreateDB()

	res, err := db.Exec("DELETE FROM connections WHERE id = ?", connection.ID)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if ra != 1 {
		return errors.New("Not one row effected")
	}

	return nil
}
