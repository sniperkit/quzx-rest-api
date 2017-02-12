package model

import "log"

type TorrentFeed struct {
	Id int `json:"id"`
	TypeId int `json:"type_id"`
	Link string `json:"link"`
	Title string `json:"title"`
	Total int `json:"total"`
	Unread int `json:"unread"`
}
type TorrentNews struct {
	Id int `json:"id"`
	FeedId int `json:"feed_id"`
	Link string `json:"link"`
	Title string `json:"title"`
	Readed int `json:"readed"`
}

func GetUnreadedTorrentFeeds() ([]*TorrentFeed, error) {

	result := []*TorrentFeed{}
	rows, err := db.Query("SELECT Id, Type_Id, Link, Title, Total, Unread FROM Feeds WHERE Unread > 0")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			f := TorrentFeed{}
			rows.Scan(&f.Id, &f.TypeId, &f.Link, &f.Title, &f.Total, &f.Unread)
			result = append(result, &f)
		}
	}

	return result, err
}

func GetUnreadedNewsByFeed(feed_id int) ([]*TorrentNews, error) {

	result := []*TorrentNews{}
	rows, err := db.Query("SELECT Id, Feed_id, Link, Title, Readed FROM News WHERE Feed_id = $1 and Readed = 0", feed_id)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			n := TorrentNews{}
			rows.Scan(&n.Id, &n.FeedId, &n.Link, &n.Title, &n.Readed)
			result = append(result, &n)
		}
	}

	return result, err
}