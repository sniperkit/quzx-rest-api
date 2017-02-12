package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/demas/cowl-services/controllers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/stack/questions/{classification}", controllers.GetStackQuestionsByClassification)
	r.HandleFunc("/api/feeds/unread", controllers.GetUnreadedTorrentFeeds)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:4000", nil)
}
