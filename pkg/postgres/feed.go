package postgres

import (
	"log"
	"strconv"

	"github.com/demas/cowl-services/pkg/quzx"
)

// represent a PostgreSQL implementation of quzx.FeedService
type FeedService struct {
}

func (s *FeedService) GetAllRssFeeds() ([]*quzx.RssFeed, error) {

	result := []*quzx.RssFeed{}
	err := db.Select(&result, "SELECT * FROM RssFeed")
	return result, err
}

func (s *FeedService) GetUnreadRssFeeds(rssType int) ([]*quzx.RssFeed, error) {

	selectQuery := `SELECT * FROM RssFeed WHERE RssType = $1 AND Unreaded > 0`
	result := []*quzx.RssFeed{}
	err := db.Select(&result, selectQuery, rssType)
	return result, err
}

func (s *FeedService) GetRssFeedById(id int) (quzx.RssFeed, error) {

	selectQuery := `SELECT * FROM RssFeed WHERE Id = $1`
	var result quzx.RssFeed
	err := db.Get(&result, selectQuery, id)
	return result, err
}

func (s *FeedService) UpdateRssFeed(feed *quzx.RssFeed) {

	tx := db.MustBegin()

	updateQuery := `UPDATE RssFeed
	                SET Link = $1, LastSyncTime = $2, AlternativeName = $3, RssType = $4,
	                    ShowContent = $5, ShowOrder = $6, Folder = $7, SyncInterval = $8
	                WHERE Id = $9"`

	_, err := tx.Exec(updateQuery,
		feed.Link,
		feed.LastSyncTime,
		feed.AlternativeName,
		feed.RssType,
		feed.ShowContent,
		feed.ShowOrder,
		feed.Folder,
		feed.SyncInterval,
		feed.Id)

	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (s *FeedService) InsertRssFeed(feed *quzx.RssFeed) {

	insertQuery := `INSERT INTO RssFeed(
				Link, SyncInterval, LastSyncTime, AlternativeName, RssType, ShowContent, ShowOrder,
			 	Folder, LimitFull, LimitHeadersOnly, Broken)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	tx := db.MustBegin()
	_, err := tx.Exec(insertQuery,
				feed.Link,
				feed.SyncInterval,
				0,
				feed.AlternativeName,
				feed.RssType,
				feed.ShowContent,
				feed.ShowOrder,
				feed.Folder,
				feed.LimitFull,
				feed.LimitHeadersOnly,
				feed.Broken)

	if err != nil {
		log.Println(err)
	}

	tx.Commit()
}

func (s *FeedService) GetRssItemsByFeedId(feed_id int) ([]*quzx.RssItem, error) {

	var feed quzx.RssFeed
	err := db.QueryRowx("SELECT ShowOrder, ShowContent, LimitFull, LimitHeadersOnly FROM RssFeed WHERE Id = $1", feed_id).StructScan(&feed)
	if err != nil {
		log.Println(err)
	}

	var strOrder string
	if feed.ShowOrder == 0 {
		strOrder = " ORDER BY Date DESC"
	} else {
		strOrder = " ORDER BY Date ASC"
	}

	var limit int
	if feed.ShowContent == 1 {
		limit = feed.LimitFull
	} else {
		limit = feed.LimitHeadersOnly
	}
	log.Println(feed)

	query := "SELECT Id, FeedId, Title, Summary, Content, Link, Date FROM RssItem WHERE FeedId = " + strconv.Itoa(feed_id) +
		" and Readed = 0" + strOrder + " LIMIT " + strconv.Itoa(limit)

	result := []*quzx.RssItem{}
	rows, err := db.Query(query)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			i := quzx.RssItem{}
			rows.Scan(&i.Id, &i.FeedId, &i.Title, &i.Summary, &i.Content, &i.Link, &i.Date)
			result = append(result, &i)
		}
	}

	return result, err
}

func (s *FeedService) SetRssItemAsReaded(id int) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE RssItem SET READED = 1 WHERE Id = $1", id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (s *FeedService) SetRssFeedAsReaded(feedId int) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE RssItem SET READED = 1 WHERE FeedId = $1", feedId)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (s *FeedService) UnsubscribeRssFeed(feedId int) {

	tx := db.MustBegin()
	_, err := tx.Exec("DELETE FROM RssItem WHERE FeedId = $1", feedId)
	if err != nil {
		log.Println(err)
	}
	_, err = tx.Exec("DELETE FROM RssFeed WHERE Id = $1", feedId)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}