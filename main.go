package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/demas/cowl-services/controllers"
)

type WithCORS struct {
	r *mux.Router
}

func (s *WithCORS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here for a Preflighted OPTIONS request.
	if req.Method == "OPTIONS" {
		return
	}

	// Lets Gorilla work
	s.r.ServeHTTP(res, req)
}

func main() {

	r := mux.NewRouter()

	// stack
	r.HandleFunc("/api/stack/tags", controllers.GetStackTags)
	r.HandleFunc("/api/stack/questions/{classification}", controllers.GetStackQuestionsByClassification)
	r.HandleFunc("/api/stack/question-as-read", controllers.SetStackQuestionAsReaded).Methods("POST")

	// torrents
	r.HandleFunc("/api/feeds/unread", controllers.GetUnreadedTorrentFeeds)
	r.HandleFunc("/api/feeds/{feed_id}/news", controllers.GetUnreadedNewsByFeed)
	r.HandleFunc("/api/news/as-read", controllers.SetTorrentNewsAsReaded).Methods("POST")

	// rss
	r.HandleFunc("/api/rss/unread", controllers.GetUnreadedRssFeeds)
	r.HandleFunc("/api/rss/{feed_id}/items", controllers.GetRssItemsByFeedId)
	r.HandleFunc("/api/rss/as-read", controllers.SetRssItemAsReaded).Methods("POST")

	http.Handle("/", &WithCORS{r})
	http.ListenAndServe("0.0.0.0:4000", nil)
}
