package repository

import "database/sql"

type AuthorizationCode struct {
	ClientId string
	Code     string
}

type Codes interface {
	Add(code AuthorizationCode) bool
	FindBy(clientId string) string
}

type PostgresCodes struct {
	Psql *sql.DB
}

func (c *PostgresCodes) Add(code AuthorizationCode) bool {
	_, err := c.Psql.Exec("insert into codes(clientId, code) values($1, $2)", code.ClientId, code.Code)

	return err == nil
}

func (c *PostgresCodes) FindBy(clientId string) string {
	row := c.Psql.QueryRow(`SELECT code FROM codes where clientId = $1`, clientId)
	var foundCode string
	row.Scan(&foundCode)

	return foundCode
}
