package postgres

import "log"
import "github.com/demas/cowl-services/quzx"

// represent a PostgreSQL implementation of quzx.TorrentService
type TorrentService struct {
}

func (s *TorrentService) GetUnreadedTorrentFeeds() ([]*quzx.TorrentFeed, error) {

	result := []*quzx.TorrentFeed{}
	rows, err := db.Query("SELECT Id, Type_Id, Link, Title, Total, Unread FROM Feeds WHERE Unread > 0")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			f := quzx.TorrentFeed{}
			rows.Scan(&f.Id, &f.TypeId, &f.Link, &f.Title, &f.Total, &f.Unread)
			result = append(result, &f)
		}
	}

	return result, err
}

func (s *TorrentService) GetUnreadedNewsByFeed(feed_id int) ([]*quzx.TorrentNews, error) {

	result := []*quzx.TorrentNews{}
	rows, err := db.Query("SELECT Id, Feed_id, Link, Title, Readed FROM News WHERE Feed_id = $1 and Readed = 0", feed_id)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			n := quzx.TorrentNews{}
			rows.Scan(&n.Id, &n.FeedId, &n.Link, &n.Title, &n.Readed)
			result = append(result, &n)
		}
	}

	return result, err
}

func (s *TorrentService) SetTorrentNewsAsReaded(news_id int) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE News SET READED = 1 WHERE Id = $1", news_id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}