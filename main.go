package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/demas/cowl-services/controllers"
	"github.com/demas/cowl-services/routers"
)

func main() {

	router := mux.NewRouter()
	router = routers.SetAuthRoute(router)

	apiRoutes := routers.InitRoutes()

	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(controllers.ValidateTokenMiddleware),
		negroni.Wrap(apiRoutes),
	))

	server := negroni.Classic()
	server.UseHandler(router)
	server.Run("0.0.0.0:4000")
}
