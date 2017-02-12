package controllers

import (
	"net/http"
	"log"
	"github.com/demas/cowl-services/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetUnreadedTorrentFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := model.GetUnreadedTorrentFeeds()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(feeds)
		w.Write(resp)
	}
}

func GetUnreadedNewsByFeed(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	feed_id, err :=  strconv.Atoi(vars["feed_id"])

	if err != nil  {
		log.Println("Error")
		w.WriteHeader(500)
	} else {
		news, err := model.GetUnreadedNewsByFeed(feed_id)

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

