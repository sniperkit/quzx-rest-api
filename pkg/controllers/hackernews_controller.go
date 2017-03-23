package controllers

import (
	"log"
	"github.com/demas/cowl-services/pkg/postgres"
	"net/http"
	"encoding/json"
)

func GetUnreadedHackerNews(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return (&postgres.HackerNewsService{}).GetUnreadedHackerNews()
}

func SetHackerNewsAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(PostData)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&postgres.HackerNewsService{}).SetHackerNewsAsReaded(bodyData.Id)
	}

	return bodyData, err
}

func SetHackerNewsAsReadedFromTime (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	type SetHackerNewsAsReadedFromTimeStruct struct {
		FromTime int64 `json:"fromTime"`
	}

	bodyData := new(SetHackerNewsAsReadedFromTimeStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		(&postgres.HackerNewsService{}).SetHackerNewsAsReadedFromTime(bodyData.FromTime)
	}

	return bodyData, err
}

func SetAllHackerNewsAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	(&postgres.HackerNewsService{}).SetAllHackerNewsAsReaded()
	return ResultOk{"ok"}, nil
}
