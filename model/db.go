package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func init() {

	database, error := sqlx.Open("postgres", "user=demas password=root host=192.168.1.71 port=5432 dbname=news sslmode=disable")
	if error != nil {
		log.Fatal("Cannot find database. Received error: " + error.Error())
	} else {
		db = database
	}
}