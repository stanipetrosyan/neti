package repository

import "database/sql"

type AuthorizationCode struct {
	ClientId string
	Code     string
}

type Codes interface {
	Add(code AuthorizationCode) bool
	FindBy(clientId string) string
	DeleteBy(clientId string) bool
}

type PostgresCodes struct {
	Psql *sql.DB
}

func (c *PostgresCodes) Add(code AuthorizationCode) bool {
	_, err := c.Psql.Exec("insert into codes(clientId, code) values($1, $2)", code.ClientId, code.Code)

	return err == nil
}

func (c *PostgresCodes) FindBy(clientId string) string {
	row := c.Psql.QueryRow(`SELECT codes FROM codes where clientId = $1`, clientId)
	var foundCode string
	row.Scan(&foundCode)

	return foundCode
}

func (c *PostgresCodes) DeleteBy(clientId string) bool {
	_, err := c.Psql.Exec(`delete from codes where clientId = $1`, clientId)

	return err == nil
}
