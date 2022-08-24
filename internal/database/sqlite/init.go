package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbname = "pet.db"
)

var DbClient *sql.DB

func init() {

	var err error

	DbClient, err = sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}

	if err = DbClient.Ping(); err == nil {
		log.Println("sqlite3 db is connected!")
	}
}
