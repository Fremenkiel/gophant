package handlers

import (
	"database/sql"
	"log"
	"time"

	"github.com/Fremenkiel/gophant/v2/internal/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type ConnectionHandler struct {
	Connection	*models.Connection
	db					*sql.DB
	cfg					pq.Config
}

func NewConnectionHandler(c *models.Connection) *ConnectionHandler {
	return &ConnectionHandler{Connection: c, cfg: pq.Config{
		Host:           c.Address,
		Port:           c.Port,
		Database: 			c.Database,
		User:           c.Username,
		Password: 			c.Password,
		ConnectTimeout: 5 * time.Second,
		SSLMode: pq.SSLModePrefer,
	}}
}

func (h *ConnectionHandler) Connect() {
	if h.db == nil {
		db, err := createConnection(h.cfg)
		if err != nil {
			log.Fatal(err)
		}

		h.db = db
	}
	err := h.db.Ping()
	if err != nil {
		db, err := createConnection(h.cfg)
		if err != nil {
			log.Fatal(err)
		}

		h.db = db
	}
	h.Connection.Status = models.ONLINE
}

func (h *ConnectionHandler) Disconnect() {
	if h.db == nil {
		return
	}
	err := h.db.Close()
	if err != nil {
		log.Fatal(err)
	}
	h.db = nil
	h.Connection.Status = models.OFFLINE
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
