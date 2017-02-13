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
	r.HandleFunc("/api/feeds/{feed_id}/news", controllers.GetUnreadedNewsByFeed)
	r.HandleFunc("/api/news/as-read", controllers.SetTorrentNewsAsReaded).Methods("POST")
	r.HandleFunc("/api/stack/question-as-read", controllers.SetStackQuestionAsReaded).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:4000", nil)
}
