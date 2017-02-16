package model

import "log"

type Reddit struct {
	Id int
	Title string
	Description string
	Link string
	AlternativeName string
	Total int
	Unreaded int
}

type RedditItem struct {
	Id int
	FeedId int
	Title string
	Summary string
	Content string
	Link string
	Date int64
}

func GetUnreadedReddits() ([]*Reddit, error) {

	result := []*Reddit{}
	rows, err := db.Query("SELECT Id, Title, Description, Link, AlternativeName, Total, Unreaded FROM Reddit WHERE Unreaded > 0")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			f := Reddit{}
			rows.Scan(&f.Id, &f.Title, &f.Description, &f.Link, &f.AlternativeName, &f.Total, &f.Unreaded)
			result = append(result, &f)
		}
	}

	return result, err
}


func GetRedditItemsByRedditId(feed_id int) ([]*RedditItem, error) {

	result := []*RedditItem{}
	rows, err := db.Query("SELECT Id, FeedId, Title, Summary, Content, Link, Date FROM RedditItem WHERE FeedId = $1 and Readed = 0 ORDER BY Date DESC", feed_id)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			i := RedditItem{}
			rows.Scan(&i.Id, &i.FeedId, &i.Title, &i.Summary, &i.Content, &i.Link, &i.Date)
			result = append(result, &i)
		}
	}

	return result, err
}

func SetRedditItemAsReaded(id int) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE RedditItem SET READED = 1 WHERE Id = $1", id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}