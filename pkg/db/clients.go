package db

import "database/sql"

type Clients interface {
	Add(client string) bool
}

type PostgresClients struct {
	Psql *sql.DB
}

func (c *PostgresClients) Add(client string) bool {
	return true
}
