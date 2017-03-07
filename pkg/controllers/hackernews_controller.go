package controllers

import (
	"log"
	"github.com/demas/cowl-services/pkg/postgres"
	"net/http"
	"encoding/json"
)

func GetUnreadedHackerNews(w http.ResponseWriter, r *http.Request) {

	news, err := (&postgres.HackerNewsService{}).GetUnreadedHackerNews()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(news)
		w.Write(resp)
	}
}

type SetHackerNewsAsReadedStruct struct {
	Id int64 `json:"id"`
}

func SetHackerNewsAsReaded (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetHackerNewsAsReadedStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		(&postgres.HackerNewsService{}).SetHackerNewsAsReaded(bodyData.Id)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

type SetHackerNewsAsReadedFromTimeStruct struct {
	FromTime int64 `json:"fromTime"`
}

func SetHackerNewsAsReadedFromTime (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetHackerNewsAsReadedFromTimeStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		(&postgres.HackerNewsService{}).SetHackerNewsAsReadedFromTime(bodyData.FromTime)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

func SetAllHackerNewsAsReaded (w http.ResponseWriter, r *http.Request) {

	(&postgres.HackerNewsService{}).SetAllHackerNewsAsReaded()
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal("{'result': 'ok'")
	w.Write(resp)
}