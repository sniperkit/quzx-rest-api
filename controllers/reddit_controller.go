package controllers

import (
	"net/http"
	"log"
	"github.com/demas/cowl-services/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetUnreadedReddits(w http.ResponseWriter, r *http.Request) {

	reddits, err := model.GetUnreadedReddits()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(reddits)
		w.Write(resp)
	}
}

func GetRedditItemsByRedditId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	feed_id, err :=  strconv.Atoi(vars["feed_id"])

	if err != nil  {
		log.Println("Error")
		w.WriteHeader(500)
	} else {
		news, err := model.GetRedditItemsByRedditId(feed_id)

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

type SetRedditItemAsReadedStruct struct {
	Id int `json:"id"`
}

func SetRedditItemAsReaded (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetRedditItemAsReadedStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		model.SetRedditItemAsReaded(bodyData.Id)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}


