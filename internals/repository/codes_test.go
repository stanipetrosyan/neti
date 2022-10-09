package repository

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

func TestPostgresCodes(t *testing.T) {
	t.Run("should add a new authorization code", func(t *testing.T) {
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

		_, err = db.Exec("CREATE TABLE codes(clientId TEXT NOT NULL, code TEXT NOT NULL)")
		if err != nil {
			log.Fatal(err)
		}

		codes := PostgresCodes{db}
		add := codes.Add(AuthorizationCode{ClientId: "aClientId", Code: "aCode"})

		assert.True(t, add)

		code := codes.FindBy("aClientId")
		assert.Equal(t, "aCode", code)
	})
}
