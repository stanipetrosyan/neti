package main

import (
	"database/sql"
	"fmt"
	"log"
	"neti/handlers"
	"neti/pkg/db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// TODO use env file
const (
	host     = "database"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "postgres"
)

func main() {
	psql := DBconnection()

	users := db.PostgresUsers{Psql: psql}
	var router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/auth", handlers.AuthApi())
	router.POST("/login", handlers.LoginApi())
	router.POST("/users", handlers.PostCreateUser(&users))

	router.Run()
}

func DBconnection() *sql.DB {
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	psql, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	err = psql.Ping()
	if err != nil {
		log.Fatal("Something went wrong with Ping", err)
	}

	return psql

}
