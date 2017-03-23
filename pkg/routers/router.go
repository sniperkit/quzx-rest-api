package routers

import (
	"github.com/gorilla/mux"
	"github.com/demas/cowl-services/pkg/controllers"
)

func InitRoutes() *mux.Router {

	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)

	// stack
	router.HandleFunc("/stack/tags", controllers.GetStackTags)
	router.HandleFunc("/stack/questions/{classification}", controllers.GetStackQuestionsByClassification)
	router.HandleFunc("/stack/question-as-read", controllers.SetStackQuestionAsReaded).Methods("POST")
	router.HandleFunc("/stack/tags/as-read", controllers.SetStackQuestionsAsReaded).Methods("POST")
	router.HandleFunc("/stack/tags/from-time/as-read", controllers.SetStackQuestionsAsReadedFromTime).Methods("POST")

	// rss
	router.HandleFunc("/rss/unread/{rss_type}", controllers.WrapHandler(controllers.GetUnreadRssFeeds))
	router.HandleFunc("/rss/allfeeds", controllers.WrapHandler(controllers.GetAllRssFeeds))
	router.HandleFunc("/rss/{feed_id}/items", controllers.WrapHandler(controllers.GetRssItemsByFeedId))
	router.HandleFunc("/rss/item/as-read", controllers.SetRssItemAsReaded).Methods("POST")
	router.HandleFunc("/rss/feed/as-read", controllers.SetRssFeedAsReaded).Methods("POST")

	router.HandleFunc("/rss/feeds/{id}", controllers.WrapHandler(controllers.GetRssFeedById)).Methods("GET")
	router.HandleFunc("/rss/feeds", controllers.PostWrapHandler(controllers.PutRssFeed)).Methods("PUT")
	router.HandleFunc("/rss/feeds", controllers.PostWrapHandler(controllers.PostRssFeed)).Methods("POST")
	router.HandleFunc("/rss/feeds/{id}", controllers.Unsubscribe).Methods("DELETE")

	// twitter
	router.HandleFunc("/twitter/favorites/{name}", controllers.GetTwitterFavourites)
	router.HandleFunc("/twitter/unfavorite", controllers.SetTwitUnfavorite).Methods("POST")

	// hacker news
	router.HandleFunc("/hn/unread", controllers.GetUnreadedHackerNews)
	router.HandleFunc("/hn/as-read", controllers.SetHackerNewsAsReaded).Methods("POST")
	router.HandleFunc("/hn/all-as-read", controllers.SetAllHackerNewsAsReaded).Methods("POST")
	router.HandleFunc("/hn/fromtime-as-read", controllers.SetHackerNewsAsReadedFromTime).Methods("POST")

	// tags
	router.HandleFunc("/tags", controllers.GetTags)
	router.HandleFunc("/tags/items/{tagId}", controllers.GetTaggedItemsByTagId).Methods("GET")
	router.HandleFunc("/tags/add-item", controllers.InsertTaggedItem).Methods("POST")
	router.HandleFunc("/tags/items/{id}", controllers.DeleteTaggedItem).Methods("DELETE")

	// bookmarks
	router.HandleFunc("/bookmarks", controllers.PostBookmark).Methods("POST")

	return router
}
