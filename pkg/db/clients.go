package db

import "database/sql"

type Clients interface {
	Add(client string) bool
	Find(id string) string
}

type PostgresClients struct {
	Psql *sql.DB
}

func (c *PostgresClients) Add(client string) bool {
	insertStmt := `insert into clients("id") values($1)`
	_, err := c.Psql.Exec(insertStmt, client)

	return err == nil
}

func (c *PostgresClients) Find(id string) string {
	row := c.Psql.QueryRow(`SELECT * FROM clients where id = $1"`, id)
	var client string
	row.Scan(&client)

	return client
}
