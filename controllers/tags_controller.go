package controllers

import (
	"net/http"
	"github.com/demas/cowl-services/model"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetTags(w http.ResponseWriter, r *http.Request) {

	tags, err := model.GetTags()

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

	items, err := model.GetTaggedItemsByTagId(tagId)

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
}

func InsertTaggedItem (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetTaggedItemStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		model.InsertTaggedItemFromStockItem(bodyData.ItemId, bodyData.TagId)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

func DeleteTaggedItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ :=  strconv.Atoi(vars["id"])

	model.DeleteTaggedItem(id)

	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal("'result':'ok'")
	w.Write(resp)
}

