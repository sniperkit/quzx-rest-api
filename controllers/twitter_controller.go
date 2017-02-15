package controllers

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/demas/cowl-services/services"
)


func GetTwitterFavourites(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	if name == ""  {
		log.Println("Attept to get twitter favourites without user name")
		w.WriteHeader(500)
	} else {
		tweets, err := services.GetFavoritesTwits(name)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		} else {
			w.Header().Add("Content-Type", "application/json")
			resp, _ := json.Marshal(tweets)
			w.Write(resp)
		}
	}
}


type SetTwitUnfavoriteStruct struct {
	Id int64 `json:"id"`
}

func SetTwitUnfavorite (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetTwitUnfavoriteStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		services.DestroyFavorites(bodyData.Id)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}