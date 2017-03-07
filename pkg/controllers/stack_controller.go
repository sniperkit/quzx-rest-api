package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/demas/cowl-services/pkg/postgres"
	"encoding/json"
)

func GetStackTags(w http.ResponseWriter, r *http.Request) {

	tags, err := (&postgres.StackService{}).GetStackTags()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(tags)
		w.Write(resp)
	}
}


func GetStackQuestionsByClassification(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	classification := vars["classification"]

	if classification == ""  {
		log.Println("Attept to get the stack questions with empty classification")
		w.WriteHeader(500)
	} else {
		questions, err := (&postgres.StackService{}).GetStackQuestionsByClassification(classification)

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

type SetStackQuestionAsReadedStruct struct {
	QuestionId int `json:"questionid"`
}

func SetStackQuestionAsReaded (w http.ResponseWriter, r *http.Request) {

	bodyData := new(SetStackQuestionAsReadedStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		(&postgres.StackService{}).SetStackQuestionAsReaded(bodyData.QuestionId)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

type UniversalPostStruct struct {
	Id int `json:"id"`
	FromTime int64 `json:"fromTime"`
	Tag string `json:"tag"`
}

func SetStackQuestionsAsReaded (w http.ResponseWriter, r *http.Request) {

	bodyData := new(UniversalPostStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		(&postgres.StackService{}).SetStackQuestionsAsReadedByClassification(bodyData.Tag)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

func SetStackQuestionsAsReadedFromTime (w http.ResponseWriter, r *http.Request) {

	bodyData := new(UniversalPostStruct)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyData)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	} else {
		(&postgres.StackService{}).SetStackQuestionsAsReadedByClassificationFromTime(bodyData.Tag, bodyData.FromTime)
		w.Header().Add("Content-Type", "application/json")
		resp, _ := json.Marshal(bodyData)
		w.Write(resp)
	}
}

