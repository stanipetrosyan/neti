package main

import (
	"database/sql"
	"fmt"
	handler "neti/internals/handler"
	repository "neti/internals/repository"
	service "neti/internals/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		FullTimestamp:          true,
	})

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psql := DBconnection()

	users := repository.PostgresUsers{Psql: psql}
	clients := repository.PostgresClients{Psql: psql}
	password := service.CryptoPassword{}
	auth := service.AuthService{}
	secret := service.CryptoSecret{}
	var router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/auth", handler.GetAuthApi())
	router.POST("/login", handler.PostLoginApi(&users, &password))
	router.POST("/users", handler.PostUsersApi(&users, &password))
	router.POST("/clients", handler.PostClientsApi(&clients, &secret))
	router.POST("/token", handler.PostTokenApi(&auth, &users, &password, &clients))
	router.GET("/authoriza", handler.GetAuthorizeApi(&clients))

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
		log.Error("Something went wrong with Ping", err)
	}

	applyMigration(psql)

	return psql

}

func applyMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != migrate.ErrNoChange {
		log.Error("Migration run failed: ", err)
	}
}
