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
