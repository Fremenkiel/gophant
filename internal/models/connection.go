package models

import (
	"time"

	"github.com/google/uuid"
)


type Connection struct {
	ID										*uuid.UUID
	Name									string
	Permission						string
	Address								string
	Port									uint16
	Database							string
	Username							string
	Password							string
	Status								Status
	DeletedAt							time.Time
}
