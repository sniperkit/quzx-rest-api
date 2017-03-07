package model

import "log"
import "github.com/demas/cowl-services/quzx"


func GetUnreadedHackerNews() ([]*quzx.HackerNews, error) {

	result := []*quzx.HackerNews{}
	rows, err := db.Query("SELECT Id, By, Score, Time, Title, Type, Url, Readed " +
		"FROM HackerNews WHERE Readed = 0 ORDER BY TIME DESC")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			n := quzx.HackerNews{}
			rows.Scan(&n.Id, &n.By, &n.Score, &n.Time, &n.Title, &n.Type, &n.Url, &n.Readed)
			result = append(result, &n)
		}
	}

	return result, err
}

func SetHackerNewsAsReaded(id int64) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE HackerNews SET READED = 1 WHERE Id = $1", id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func SetHackerNewsAsReadedFromTime(t int64) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE HackerNews SET READED = 1 WHERE Time < $1", t)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func SetAllHackerNewsAsReaded() {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE HackerNews SET READED = 1")
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

