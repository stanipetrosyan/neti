package main

import (
	"database/sql"
	"fmt"
	"log"
	"neti/handlers"
	"neti/pkg/crypto"
	"neti/pkg/db"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psql := DBconnection()

	users := db.PostgresUsers{Psql: psql}
	clients := db.PostgresClients{Psql: psql}
	password := crypto.CryptoPassword{}
	var router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/auth", handlers.AuthApi())
	router.POST("/login", handlers.LoginApi(&users, &password))
	router.POST("/users", handlers.PostCreateUser(&users, &password))
	router.POST("/clients", handlers.PostClientsApi(&clients))

	router.Run()
}

func DBconnection() *sql.DB {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	psql, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	err = psql.Ping()
	if err != nil {
		log.Fatal("Something went wrong with Ping", err)
	}

	applyMigration(psql)

	return psql

}

func applyMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://pkg/db/migrations/", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Println("si spacca qui", err)
	}
}
