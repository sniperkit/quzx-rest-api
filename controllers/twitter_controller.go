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