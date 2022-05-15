package repository

import (
	"database/sql"
	"neti/internals/domain"
)

type Clients interface {
	Add(client domain.Client) bool
	FindBy(id string) domain.Client
	Exist(id string) bool
}

type PostgresClients struct {
	Psql *sql.DB
}

func (c *PostgresClients) Add(client domain.Client) bool {
	insertStmt := `insert into clients values($1, $2)`
	_, err := c.Psql.Exec(insertStmt, client.ClientId, client.ClientSecret)

	return err == nil
}

func (c *PostgresClients) FindBy(id string) domain.Client {
	row := c.Psql.QueryRow(`SELECT * FROM clients where clientId = $1`, id)
	var clientId string
	var clientSecret string
	err := row.Scan(&clientId, &clientSecret)

	if err != nil {
		return domain.Client{}
	}

	return domain.Client{ClientId: clientId, ClientSecret: clientSecret}
}

func (c *PostgresClients) Exist(id string) bool {
	return c.FindBy(id).ClientId == id
}
