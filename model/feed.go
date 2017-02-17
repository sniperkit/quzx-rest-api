package model

import "log"

type RssFeed struct {
	Id int
	Title string
	Description string
	Link string
	ImageUrl string
	AlternativeName string
	Total int
	Unreaded int
	RssType int
}

type RssItem struct {
	Id int
	FeedId int
	Title string
	Summary string
	Content string
	Link string
	Date int64
}

func GetUnreadedRssFeeds(rssType int) ([]*RssFeed, error) {

	result := []*RssFeed{}
	rows, err := db.Query("SELECT Id, Title, Description, Link, ImageUrl, AlternativeName, Total, Unreaded FROM RssFeed WHERE RssType = $1 AND Unreaded > 0", rssType)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			f := RssFeed{}
			rows.Scan(&f.Id, &f.Title, &f.Description, &f.Link, &f.ImageUrl, &f.AlternativeName, &f.Total, &f.Unreaded)
			result = append(result, &f)
		}
	}

	return result, err
}


func GetRssItemsByFeedId(feed_id int) ([]*RssItem, error) {

	result := []*RssItem{}
	rows, err := db.Query("SELECT Id, FeedId, Title, Summary, Content, Link, Date FROM RssItem WHERE FeedId = $1 and Readed = 0 ORDER BY Date DESC", feed_id)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			i := RssItem{}
			rows.Scan(&i.Id, &i.FeedId, &i.Title, &i.Summary, &i.Content, &i.Link, &i.Date)
			result = append(result, &i)
		}
	}

	return result, err
}

func SetRssItemAsReaded(id int) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE RssItem SET READED = 1 WHERE Id = $1", id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}