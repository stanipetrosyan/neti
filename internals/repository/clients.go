package repository

import (
	"database/sql"
	"neti/internals/domain"
)

type Clients interface {
	Add(client domain.Client) bool
	FindBy(id string) domain.Client
}

type PostgresClients struct {
	Psql *sql.DB
}

func (c *PostgresClients) Add(client domain.Client) bool {
	insertStmt := `insert into clients("id") values($1)`
	_, err := c.Psql.Exec(insertStmt, client.ClientId)

	return err == nil
}

func (c *PostgresClients) FindBy(id string) domain.Client {
	row := c.Psql.QueryRow(`SELECT * FROM clients where id = $1`, id)
	var clientId string
	row.Scan(&clientId)

	return domain.Client{ClientId: clientId}
}
