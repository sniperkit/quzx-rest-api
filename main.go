package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/demas/cowl-services/controllers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/stack/questions/{classification}", controllers.GetStackQuestionsByClassification)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:4000", nil)
}
