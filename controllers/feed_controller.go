package controllers

import (
	"net/http"
	"log"
	"github.com/demas/cowl-services/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetRssFeeds(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rss_type, err :=  strconv.Atoi(vars["rss_type"])

	var onlyUnreaded bool
	if vars["only_unreaded"] == "1" {
		onlyUnreaded = true
	} else {
		onlyUnreaded = false
	}

	feeds, err := model.GetRssFeeds(rss_type, onlyUnreaded)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(feeds)
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
		news, err := model.GetRssItemsByFeedId(feed_id)

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
		model.SetRssItemAsReaded(bodyData.Id)
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
		model.SetRssFeedAsReaded(bodyData.FeedId)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}


