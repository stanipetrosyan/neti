package db

type Users interface {
	Add(username string, password string) bool
}

type PostgresUsers struct {
}

func (u *PostgresUsers) Add(username string, password string) bool {
	return true
}
