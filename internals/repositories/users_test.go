package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"neti/internals/domain"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

func TestPostgresUsers(t *testing.T) {
	t.Run("should add a user", func(t *testing.T) {

		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatal(err)
		}

		resource, _ := pool.Run("postgres", "13", []string{"POSTGRES_PASSWORD=password", "POSTGRES_USER=user"})
		defer pool.Purge(resource)

		connection := fmt.Sprintf("host=localhost port=%s user=user password=password dbname=postgres sslmode=disable", resource.GetPort("5432/tcp"))
		var db *sql.DB

		if err = pool.Retry(func() error {
			db, err = sql.Open("postgres", connection)
			return db.Ping()
		}); err != nil {
			log.Fatal(err)
		}

		// migration db
		_, err = db.Exec("CREATE TABLE users(username text, password text)")
		if err != nil {
			log.Fatal(err)
		}

		clients := PostgresUsers{db}
		clients.Add(domain.User{Username: "user", Password: "pass"})

		username, password := clients.FindBy("user")
		assert.Equal(t, username, "user")
		assert.Equal(t, password, "pass")

	})
}
