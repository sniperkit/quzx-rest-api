package model

import "log"

type HackerNews struct {
	Id int64
	By string
	Score int
	Time int64
	Title string
	Type string
	Url string
	Readed int
}


func GetUnreadedHackerNews() ([]*HackerNews, error) {

	result := []*HackerNews{}
	rows, err := db.Query("SELECT Id, By, Score, Time, Title, Type, Url, Readed " +
		"FROM HackerNews WHERE Readed = 0 ORDER BY TIME DESC")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			n := HackerNews{}
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

