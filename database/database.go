package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var connStr string = "user=root password=root dbname=database sslmode=disable"

var Conn *sql.DB = ConnDB(connStr)

func ConnDB(connStr string) *sql.DB {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
