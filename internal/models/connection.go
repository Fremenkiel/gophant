package models

import "github.com/google/uuid"


type Connection struct {
	ID										*uuid.UUID
	Name									string
	Address								string
	Port									uint16
	Database							string
	Username							string
	Password							string
	Status								Status
}
