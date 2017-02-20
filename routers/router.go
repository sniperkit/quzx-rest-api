package routers

import (
	"github.com/gorilla/mux"
	"github.com/demas/cowl-services/controllers"
)

func InitRoutes() *mux.Router {

	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)

	// stack
	router.HandleFunc("/stack/tags", controllers.GetStackTags)
	router.HandleFunc("/stack/questions/{classification}", controllers.GetStackQuestionsByClassification)
	router.HandleFunc("/stack/question-as-read", controllers.SetStackQuestionAsReaded).Methods("POST")

	// torrents
	router.HandleFunc("/feeds/unread", controllers.GetUnreadedTorrentFeeds)
	router.HandleFunc("/feeds/{feed_id}/news", controllers.GetUnreadedNewsByFeed)
	router.HandleFunc("/news/as-read", controllers.SetTorrentNewsAsReaded).Methods("POST")

	// rss
	router.HandleFunc("/rss/unread/{rss_type}", controllers.GetUnreadedRssFeeds)
	router.HandleFunc("/rss/{feed_id}/items", controllers.GetRssItemsByFeedId)
	router.HandleFunc("/rss/as-read", controllers.SetRssItemAsReaded).Methods("POST")

	// twitter
	router.HandleFunc("/twitter/favorites/{name}", controllers.GetTwitterFavourites)
	router.HandleFunc("/twitter/unfavorite", controllers.SetTwitUnfavorite).Methods("POST")

	// hacker news
	router.HandleFunc("/hn/unread", controllers.GetUnreadedHackerNews)
	router.HandleFunc("/hn/as-read", controllers.SetHackerNewsAsReaded).Methods("POST")

	return router
}
