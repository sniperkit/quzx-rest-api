package controllers

import (
	"net/http"
	"github.com/demas/cowl-services/pkg/postgres"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetTags(w http.ResponseWriter, r *http.Request) {

	tags, err := (&postgres.TagsService{}).GetTags()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(tags)
		w.Write(resp)
	}
}

func GetTaggedItemsByTagId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tagId, err :=  strconv.Atoi(vars["tagId"])
	items, err := (&postgres.TagsService{}).GetTaggedItemsByTagId(tagId)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(items)
		w.Write(resp)
	}
}

type SetTaggedItemStruct struct {
	ItemId int `json:"itemId"`
	TagId int `json:"tagId"`
	Source int `json:"source"`
}

func InsertTaggedItem (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetTaggedItemStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		if bodyData.Source == 1 {
			(&postgres.TagsService{}).InsertTaggedItemFromStockItem(bodyData.ItemId, bodyData.TagId)
		} else if bodyData.Source == 2 {
			(&postgres.TagsService{}).InsertTaggedItemFromRss(bodyData.ItemId, bodyData.TagId)
		}

		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

func DeleteTaggedItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ :=  strconv.Atoi(vars["id"])

	(&postgres.TagsService{}).DeleteTaggedItem(id)

	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal("'result':'ok'")
	w.Write(resp)
}

