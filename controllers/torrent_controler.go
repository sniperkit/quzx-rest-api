package controllers

import (
	"net/http"
	"log"
	"github.com/demas/cowl-services/postgres"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetUnreadedTorrentFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := (&postgres.TorrentService{}).GetUnreadedTorrentFeeds()

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
		news, err := (&postgres.TorrentService{}).GetUnreadedNewsByFeed(feed_id)

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

type SetTorrentAsReaded struct {
	Id int `json:"id"`
}

func SetTorrentNewsAsReaded (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetTorrentAsReaded)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		(&postgres.TorrentService{}).SetTorrentNewsAsReaded(bodyData.Id)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}


