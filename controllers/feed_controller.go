package controllers

import (
	"net/http"
	"log"
	"github.com/demas/cowl-services/postgres"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/demas/cowl-services/quzx"
)

func GetUnreadRssFeeds(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rss_type, err :=  strconv.Atoi(vars["rss_type"])

	feeds, err := postgres.GetUnreadRssFeeds(rss_type)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(feeds)
		w.Write(resp)
	}
}

func GetAllRssFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := postgres.GetAllRssFeeds()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(feeds)
		w.Write(resp)
	}
}

func GetRssFeedById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err :=  strconv.Atoi(vars["id"])

	feed, err := postgres.GetRssFeedById(id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(feed)
		w.Write(resp)
	}
}

func PutRssFeed(w http.ResponseWriter, r *http.Request) {

	bodyData := new(quzx.RssFeed)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		postgres.UpdateRssFeed(bodyData)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

func PostRssFeed(w http.ResponseWriter, r *http.Request) {

	bodyData := new(quzx.RssFeed)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		postgres.InsertRssFeed(bodyData)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

func GetRssItemsByFeedId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	feed_id, err :=  strconv.Atoi(vars["feed_id"])

	if err != nil  {
		log.Println("Error")
		w.WriteHeader(500)
	} else {
		news, err := postgres.GetRssItemsByFeedId(feed_id)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		} else {
			w.Header().Add("Content-Type", "application/json")
			resp, _ := json.Marshal(news)
			w.Write(resp)
		}

	}
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
		postgres.SetRssItemAsReaded(bodyData.Id)
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
		postgres.SetRssFeedAsReaded(bodyData.FeedId)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

func Unsubscribe (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	feedid, _ :=  strconv.Atoi(vars["id"])

	postgres.UnsubscribeRssFeed(feedid)

	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal("'result':'ok'")
	w.Write(resp)
}



