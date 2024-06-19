package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	connStr := "user=postgres password=123456 dbname=postgres sslmode=disable"
	var db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("Conectou ao banco!\n")
	}

	return db
}
