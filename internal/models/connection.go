package models

import "github.com/google/uuid"


type Connection struct {
	ID										uuid.UUID
	Name									string
	ConnectionString			string
	Status								Status
}
