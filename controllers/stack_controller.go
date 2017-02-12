package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/demas/cowl-services/model"
	"encoding/json"
)

func GetStackQuestionsByClassification(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	classification := vars["classification"]

	if classification == ""  {
		log.Println("Attept to get the stack questions with empty classification")
		w.WriteHeader(500)
	} else {
		questions, err := model.GetStackQuestionsByClassification(classification)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		} else {
			w.Header().Add("Content-Type", "application/json")
			resp, _ := json.Marshal(questions)
			w.Write(resp)
		}

	}
}
