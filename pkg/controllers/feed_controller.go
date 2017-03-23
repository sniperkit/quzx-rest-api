package controllers

import (
	"net/http"
	"github.com/demas/cowl-services/pkg/postgres"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/demas/cowl-services/pkg/quzx"

)

func GetUnreadRssFeeds(w http.ResponseWriter, r *http.Request) (interface{}, error)  {

	rss_type, _ :=  strconv.Atoi(mux.Vars(r)["rss_type"])
	return (&postgres.FeedService{}).GetUnreadRssFeeds(rss_type)
}

func GetAllRssFeeds(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&postgres.FeedService{}).GetAllRssFeeds()
}

func GetRssFeedById(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	id, _ :=  strconv.Atoi(mux.Vars(r)["id"])
	return (&postgres.FeedService{}).GetRssFeedById(id)
}

func GetRssItemsByFeedId(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	var news []*quzx.RssItem
	feed_id, err :=  strconv.Atoi(mux.Vars(r)["feed_id"])
	if err == nil {
		news, err = (&postgres.FeedService{}).GetRssItemsByFeedId(feed_id)
	}
	return news, err
}

func PutRssFeed(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(quzx.RssFeed)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&postgres.FeedService{}).UpdateRssFeed(bodyData)
	}

	return bodyData, err
}

func PostRssFeed(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(quzx.RssFeed)
	err := json.NewDecoder(r.Body).Decode(&bodyData)

	if err == nil {
		(&postgres.FeedService{}).InsertRssFeed(bodyData)
	}

	return bodyData, err
}

func SetRssItemAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(PostData)
	err := json.NewDecoder(r.Body).Decode(&bodyData)

	if err == nil {
		(&postgres.FeedService{}).SetRssItemAsReaded(int(bodyData.Id))
	}

	return bodyData, err
}

func SetRssFeedAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	type SetRssFeedAsReadedStruct struct {
		FeedId int `json:"feed_id"`
	}

	bodyData := new(SetRssFeedAsReadedStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)

	if err == nil {
		(&postgres.FeedService{}).SetRssFeedAsReaded(bodyData.FeedId)
	}

	return bodyData, err
}

func Unsubscribe (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	feedid, err :=  strconv.Atoi(mux.Vars(r)["id"])
	(&postgres.FeedService{}).UnsubscribeRssFeed(feedid)
	return ResultOk{"ok"}, err
}



