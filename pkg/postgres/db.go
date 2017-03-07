package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"fmt"
)

type DbOptions struct {
	user string
	password string
	host string
	port string
	dbname string
}

var db *sqlx.DB

func open(ops *DbOptions) (*sqlx.DB, error) {

	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		ops.user, ops.password, ops.host, ops.port, ops.dbname)
	database, error := sqlx.Open("postgres", connString)
	if error != nil {
		log.Fatal("Cannot find database. Received error: " + error.Error())
		return nil, error
	} else {
		return database, nil
	}
}

func init() {

	var ops = DbOptions{os.Getenv("DBUSER"), os.Getenv("DBPASS"),
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME")}

	database, error := open(&ops)
	if error != nil {
		panic(error)
	} else {
		db = database
	}
}