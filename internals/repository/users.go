package repository

import (
	"database/sql"
	"neti/internals/domain"
)

type Users interface {
	Add(user domain.User) bool
	FindBy(username string) (string, string)
	AddRole(username string, role string) bool
}

type PostgresUsers struct {
	Psql *sql.DB
}

func (u *PostgresUsers) Add(user domain.User) bool {
	insertStmt := `insert into users("username", "password") values($1, $2)`
	_, err := u.Psql.Exec(insertStmt, user.Username, user.Password)

	return err == nil
}

func (u *PostgresUsers) FindBy(username string) (string, string) {
	row := u.Psql.QueryRow(`SELECT * FROM users where username = $1`, username)
	var foundUsername string
	var foundPassword string
	row.Scan(&foundUsername, &foundPassword)

	return foundUsername, foundPassword
}

func (u *PostgresUsers) AddRole(username string, role string) bool {
	return false
}
