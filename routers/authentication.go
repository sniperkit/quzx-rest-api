package routers

import (
	"github.com/gorilla/mux"
	"github.com/demas/cowl-services/controllers"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")

	return router
}
