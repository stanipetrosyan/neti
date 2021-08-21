package db

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
			log.Fatal("error")
		}

		resource, _ := pool.Run("postgres", "13", []string{"POSTGRES_PASSWORD=password", "POSTGRES_USER=user"})
		println(resource.Container.Config)
		println(resource.GetPort("5432/tcp"))

		connection := fmt.Sprintf("host=localhost port=%s user=user password=password dbname=postgres sslmode=disable", resource.GetPort("5432/tcp"))

		if err = pool.Retry(func() error {
			db, err := sql.Open("postgres", connection)
			if err != nil {
				log.Fatal(err)
			}
			err = db.Ping()
			return err
		}); err != nil {
			log.Fatal("err", err)
		}

		assert.Equal(t, resource.GetPort("5432/tcp"), "5432")
		assert.Equal(t, resource.GetPort("5432/tcp"), "5432")
	})
}
