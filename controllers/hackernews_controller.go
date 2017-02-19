package controllers

import (
	"log"
	"github.com/demas/cowl-services/model"
	"net/http"
	"encoding/json"
)

func GetUnreadedHackerNews(w http.ResponseWriter, r *http.Request) {

	news, err := model.GetUnreadedHackerNews()

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
		model.SetHackerNewsAsReaded(bodyData.Id)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}
