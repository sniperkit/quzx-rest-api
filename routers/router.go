package routers

import (
	"github.com/gorilla/mux"
	"github.com/demas/cowl-services/controllers"
	"github.com/urfave/negroni"
	"net/http"
)

func InitRoutes() *mux.Router {

	router := mux.NewRouter()

	router = SetAuthenticationRoutes(router)

	// stack
	router.HandleFunc("/api/stack/tags", controllers.GetStackTags)
	router.HandleFunc("/api/stack/questions/{classification}", controllers.GetStackQuestionsByClassification)
	router.HandleFunc("/api/stack/question-as-read", controllers.SetStackQuestionAsReaded).Methods("POST")

	// torrents
	router.Handle("/api/feeds/unread", negroni.New(
		negroni.HandlerFunc(controllers.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(controllers.GetUnreadedTorrentFeeds)),
	))

	router.HandleFunc("/api/feeds/{feed_id}/news", controllers.GetUnreadedNewsByFeed)
	router.HandleFunc("/api/news/as-read", controllers.SetTorrentNewsAsReaded).Methods("POST")

	// rss
	router.HandleFunc("/api/rss/unread/{rss_type}", controllers.GetUnreadedRssFeeds)
	router.HandleFunc("/api/rss/{feed_id}/items", controllers.GetRssItemsByFeedId)
	router.HandleFunc("/api/rss/as-read", controllers.SetRssItemAsReaded).Methods("POST")

	// twitter
	router.HandleFunc("/api/twitter/favorites/{name}", controllers.GetTwitterFavourites)
	router.HandleFunc("/api/twitter/unfavorite", controllers.SetTwitUnfavorite).Methods("POST")

	// hacker news
	router.HandleFunc("/api/hn/unread", controllers.GetUnreadedHackerNews)
	router.HandleFunc("/api/hn/as-read", controllers.SetHackerNewsAsReaded).Methods("POST")

	return router
}
