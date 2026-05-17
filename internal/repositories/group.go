package repositories

import (
	"errors"

	"github.com/Fremenkiel/gophant/v2/internal/database"
	"github.com/Fremenkiel/gophant/v2/internal/models"
)

type GroupRepository struct {
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{}
}

func (r *GroupRepository) GetAll() ([]models.Group, error) {
	db := database.CurrentDB()

	rows, err := db.Query("SELECT * FROM groups")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group

	for rows.Next() {
		var c models.Group
		if err := rows.Scan(&c.ID, &c.Name, &c.R, &c.G, &c.B); err != nil {
			return nil, err
		}
		groups = append(groups, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return groups, nil
}

func (r *GroupRepository) Get(id uint) (models.Group, error) {
	db := database.CurrentDB()

	rows, err := db.Query("SELECT * FROM groups WHERE id = ?")
	if err != nil {
		return models.Group{}, err
	}
	defer rows.Close()

	var g models.Group

	for rows.Next() {
		if err := rows.Scan(&g.ID, &g.Name, &g.R, &g.G, &g.B); err != nil {
			return models.Group{}, err
		}
	}
	if err = rows.Err(); err != nil {
		return models.Group{}, err
	}

	return g, nil
}

func (r *GroupRepository) Create(group models.Group) error {
	db := database.CurrentDB()

	stmt, err := db.Prepare("INSERT INTO groups (name, R, G, B) VALUES (?, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(group.Name, group.R, group.G, group.B)
	stmt.Close()
	return err
}

func (r *GroupRepository) Update(group models.Group) error {
	db := database.CurrentDB()

	stmt, err := db.Prepare("UPDATE groups SET name = ?, R = ?, G = ?, B = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(group.Name, group.R, group.G, group.B, group.ID)
	stmt.Close()
	return err
}

func (r *GroupRepository) Delete(group models.Group) error {
	db := database.CurrentDB()

	res, err := db.Exec("DELETE FROM groups WHERE id = ?", group.ID)
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
