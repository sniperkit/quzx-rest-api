package postgres

import "log"
import "github.com/demas/cowl-services/pkg/quzx"

// represent a PostgreSQL implementation of quzx.HackerNewsService
type HackerNewsService struct {
}

func (s *HackerNewsService) GetUnreadedHackerNews() ([]*quzx.HackerNews, error) {

	result := []*quzx.HackerNews{}
	err := db.Select(&result, "SELECT * FROM HackerNews WHERE Readed = 0 ORDER BY TIME DESC")
	return result, err
}

func (s *HackerNewsService) SetHackerNewsAsReaded(id int64) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE HackerNews SET READED = 1 WHERE Id = $1", id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (s *HackerNewsService) SetHackerNewsAsReadedFromTime(t int64) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE HackerNews SET READED = 1 WHERE Time < $1", t)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (s *HackerNewsService) SetAllHackerNewsAsReaded() {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE HackerNews SET READED = 1")
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

