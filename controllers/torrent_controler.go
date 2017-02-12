package controllers

import (
	"net/http"
	"log"
	"github.com/demas/cowl-services/model"
	"encoding/json"
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
