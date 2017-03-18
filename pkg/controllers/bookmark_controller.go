package controllers

import (
	"net/http"
	"encoding/json"
	"log"
)

func PostBookmark(w http.ResponseWriter, r *http.Request) {

	type PostBookmark struct {
		bookmark map[string]string `json:"bookmark"`
	}

	bodyData := new(PostBookmark)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		//(&postgres.BookmarkRepository{}).InsertBookmark(bodyData.bookmark, bodyData.tags)
		log.Println(bodyData.bookmark["title"])
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}
