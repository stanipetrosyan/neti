package repository

import (
	"database/sql"
	"neti/internals/domain"
)

type Users interface {
	Add(user domain.User) bool
	FindBy(username string) domain.User
	AddRole(username string, role string) bool
}

type PostgresUsers struct {
	Psql *sql.DB
}

func (u *PostgresUsers) Add(user domain.User) bool {
	insertStmt := `insert into users("username", "password", "role") values($1, $2, $3)`
	_, err := u.Psql.Exec(insertStmt, user.Username, user.Password, user.Role)

	return err == nil
}

func (u *PostgresUsers) FindBy(username string) domain.User {
	row := u.Psql.QueryRow(`SELECT * FROM users where username = $1`, username)

	var user domain.User
	row.Scan(&user.Username, &user.Password, &user.Role)

	return user
}

func (u *PostgresUsers) AddRole(username string, role string) bool {
	updateStmt := `update users set role = $1 where username = $2`
	_, err := u.Psql.Exec(updateStmt, role, username)

	return err == nil

}
