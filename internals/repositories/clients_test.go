package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

func TestPostgresClients(t *testing.T) {

	t.Run("should add a new client", func(t *testing.T) {
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
		_, err = db.Exec("CREATE TABLE clients(id text)")
		if err != nil {
			log.Fatal(err)
		}

		clients := PostgresClients{db}
		clients.Add("aClient")

		client := clients.Find("aClient")
		assert.Equal(t, client, "aClient")
	})
}
