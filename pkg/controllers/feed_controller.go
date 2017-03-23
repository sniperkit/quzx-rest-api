package controllers

import (
	"net/http"
	"log"
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

type SetRssItemAsReadedStruct struct {
	Id int `json:"id"`
}

func SetRssItemAsReaded (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetRssItemAsReadedStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		(&postgres.FeedService{}).SetRssItemAsReaded(bodyData.Id)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

type SetRssFeedAsReadedStruct struct {
	FeedId int `json:"feed_id"`
}

func SetRssFeedAsReaded (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetRssFeedAsReadedStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		(&postgres.FeedService{}).SetRssFeedAsReaded(bodyData.FeedId)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

func Unsubscribe (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	feedid, _ :=  strconv.Atoi(vars["id"])

	(&postgres.FeedService{}).UnsubscribeRssFeed(feedid)

	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal("'result':'ok'")
	w.Write(resp)
}



