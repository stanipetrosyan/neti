package repository

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
	_, err = db.Exec("CREATE TABLE users(username text, password text, role text)")
	if err != nil {
		log.Fatal(err)
	}
	users := PostgresUsers{db}

	t.Run("should add a user", func(t *testing.T) {
		users.Add(domain.User{Username: "user", Password: "pass"})

		user := users.FindBy("user")
		println(user.Username)
		assert.Equal(t, user.Username, "user")
		assert.Equal(t, user.Password, "pass")
		assert.Equal(t, user.Role, "")
	})

	t.Run("should add a role to user", func(t *testing.T) {
		users.Add(domain.User{Username: "user", Password: "pass"})
		users.AddRole("user", "aRole")

		user := users.FindBy("user")
		assert.Equal(t, "user", user.Username)
		assert.Equal(t, "pass", user.Password)
		assert.Equal(t, "aRole", user.Role)
	})
}
