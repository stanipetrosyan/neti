package db

import "database/sql"

type Clients interface {
	Add(client string) bool
}

type PostgresClients struct {
	Psql *sql.DB
}

func (c *PostgresClients) Add(client string) bool {
	insertStmt := `insert into clients("id") values($1)`
	_, err := c.Psql.Exec(insertStmt, client)

	return err == nil
}
