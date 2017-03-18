package postgres

import (
	"log"
	"fmt"
	"strconv"

	"github.com/demas/cowl-services/pkg/quzx"
)

// represent a PostgreSQL implementation of quzx.FeedService
type FeedService struct {

}

func (s *FeedService) GetAllRssFeeds() ([]*quzx.RssFeed, error) {

	result := []*quzx.RssFeed{}
	rows, err := db.Query("SELECT Id, Title, Description, Link, LastSyncTime, ImageUrl, " +
		                     "AlternativeName, Total, Unreaded, " +
		                     "RssType, ShowContent, ShowOrder, Folder, Broken " +
		                     "FROM RssFeed")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			f := quzx.RssFeed{}
			rows.Scan(&f.Id, &f.Title, &f.Description, &f.Link, &f.LastSyncTime,
				&f.ImageUrl, &f.AlternativeName, &f.Total,
				&f.Unreaded, &f.RssType, &f.ShowContent, &f.ShowOrder, &f.Folder, &f.Broken)
			result = append(result, &f)
		}
	}

	return result, err
}


func (s *FeedService) GetUnreadRssFeeds(rssType int) ([]*quzx.RssFeed, error) {

	result := []*quzx.RssFeed{}
	query := fmt.Sprintf("SELECT Id, Title, Description, Link, LastSyncTime, ImageUrl, " +
				              "AlternativeName, Total, Unreaded, SyncInterval, " +
		                              "RssType, ShowContent, ShowOrder, Folder " +
		                              "FROM RssFeed WHERE RssType = %d AND Unreaded > 0", rssType)

	rows, err := db.Query(query)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			f := quzx.RssFeed{}
			rows.Scan(&f.Id, &f.Title, &f.Description, &f.Link, &f.LastSyncTime,
				&f.ImageUrl, &f.AlternativeName, &f.Total,
				&f.Unreaded, &f.SyncInterval, &f.RssType, &f.ShowContent, &f.ShowOrder, &f.Folder)
			result = append(result, &f)
		}
	}

	return result, err
}


func (s *FeedService) GetRssFeedById(id int) (quzx.RssFeed, error) {

	var result quzx.RssFeed
	query := fmt.Sprintf("SELECT Id, Title, Description, Link, LastSyncTime, ImageUrl, " +
		                     "AlternativeName, Total, Unreaded, SyncInterval, RssType, ShowContent, ShowOrder, Folder " +
		                     "FROM RssFeed WHERE Id = %d", id)
	err := db.Get(&result, query)

	if err != nil {
		log.Println(err)
	}

	return result, err
}

func (s *FeedService) UpdateRssFeed(feed *quzx.RssFeed) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE RssFeed SET Link = $1, LastSyncTime = $2, AlternativeName = $3, " +
				 "RssType = $4, ShowContent = $5, ShowOrder = $6, Folder = $7, SyncInterval = $8 WHERE Id = $9",
				 feed.Link, feed.LastSyncTime, feed.AlternativeName, feed.RssType, feed.ShowContent,
				 feed.ShowOrder, feed.Folder, feed.SyncInterval, feed.Id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (s *FeedService) InsertRssFeed(feed *quzx.RssFeed) {

	insertQuery := `INSERT INTO RssFeed
			(Link, SyncInterval, LastSyncTime, AlternativeName, RssType, ShowContent, ShowOrder,
			 Folder, LimitFull, LimitHeadersOnly, Broken)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	tx := db.MustBegin()
	_, err := tx.Exec(insertQuery,
				feed.Link,
				feed.SyncInterval, 0,
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