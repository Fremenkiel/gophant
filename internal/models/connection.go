package models

type Connection struct {
	ID										uint
	Name									string
	Permission						string
	Address								string
	Port									uint16
	Database							string
	Username							string
	Password							string
	Status								Status
}
