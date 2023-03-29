package main

import (
	"challenge-06/controllers"
	"challenge-06/routers"
	"database/sql"
	"fmt"

	// initial Database framework
	_ "github.com/lib/pq"
)

/*
go get -u github.com/gin-gonic/gin
go get -u github.com/lib/pq
*/
const (
	host     = "localhost"
	dbport   = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "db-go-sql"
	PORT     = ":8080"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, dbport, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	controllers.DB = db

	routers.StartServer().Run(PORT)
	fmt.Println("Successfully connected to database")
}
