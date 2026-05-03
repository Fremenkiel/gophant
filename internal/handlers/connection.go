package handlers

import (
	"database/sql"
	"log"
	"time"

	"github.com/Fremenkiel/gophant/v2/internal/interfaces"
	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type ConnectionHandler struct {
	Connection	*models.Connection

	c					*sql.DB
	cfg					pq.Config
	db					[]models.Database
	reporter		interfaces.ErrorReporter
}

func NewConnectionHandler(r interfaces.ErrorReporter, c *models.Connection) *ConnectionHandler {
	return &ConnectionHandler{Connection: c, cfg: pq.Config{
		Host:           c.Address,
		Port:           c.Port,
		Database: 			c.Database,
		User:           c.Username,
		Password: 			c.Password,
		ConnectTimeout: 5 * time.Second,
		SSLMode: pq.SSLModePrefer,
	}, reporter: r}
}

func (h *ConnectionHandler) Connect() {
	if h.c == nil {
		db, err := createConnection(h.cfg)
		if err != nil {
			h.reporter.Report(err)
			return
		}

		h.c = db
	}
	err := h.c.Ping()
	if err != nil {
		db, err := createConnection(h.cfg)
		if err != nil {
			h.reporter.Report(err)
			return
		}

		h.c = db
	}
	h.Connection.Status = models.ONLINE
}

func (h *ConnectionHandler) Disconnect() {
	if h.c == nil {
		return
	}
	err := h.c.Close()
	if err != nil {
			h.reporter.Report(err)
			return
	}
	h.c = nil
	h.Connection.Status = models.OFFLINE
}

func (h *ConnectionHandler) GetDatabases(refresh func()) {
	if h.c == nil || h.Connection.Status != models.ONLINE {
		h.Connect()
		refresh()
	}

	rows, err := h.c.Query("SELECT datname FROM pg_database;")
	if err != nil {
			h.reporter.Report(err)
			return
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var databases []models.Database

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var db models.Database
		if err := rows.Scan(&db.Name); err != nil {
			h.reporter.Report(err)
			return
		}
		databases = append(databases, db)
	}
	if err = rows.Err(); err != nil {
			h.reporter.Report(err)
			return
	}
	h.db = databases
	log.Print(databases)
}

func createConnection(cfg pq.Config) (*sql.DB, error) {
	cf, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(cf)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
