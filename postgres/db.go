package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sqlx.DB

func init() {

	datasourceName := "user=" + os.Getenv("DBUSER") + " password=" + os.Getenv("DBPASS") +
		" host=" + os.Getenv("DBHOST") + " port=" + os.Getenv("DBPORT")+
		" dbname=" + os.Getenv("DBNAME") + " sslmode=disable"
	database, error := sqlx.Open("postgres", datasourceName)
	if error != nil {
		log.Fatal("Cannot find database. Received error: " + error.Error())
	} else {
		db = database
	}
}