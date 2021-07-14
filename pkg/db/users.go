package db

import "database/sql"

// Add User model struct
type Users interface {
	Add(username string, password string) bool
	FindUserBy(username string) (string, string)
}

type PostgresUsers struct {
	Psql *sql.DB
}

// Add tests
func (u *PostgresUsers) Add(username string, password string) bool {
	insertStmt := `insert into "users"("username", "password") values($1, $2)`
	_, err := u.Psql.Exec(insertStmt, username, password)

	return err == nil
}

func (u *PostgresUsers) FindUserBy(username string) (string, string) {
	row := u.Psql.QueryRow(`SELECT * FROM "users where username = $1"`, username)
	var foundUsername string
	var foundPassword string
	row.Scan(&foundUsername, &foundPassword)

	return foundUsername, foundPassword
}
