package controllers

import (
	"net/http"
	"log"
	"encoding/json"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request) (interface{}, error)

func WrapHandler(handler HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		data, err := handler(w, req)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		} else {
			w.Header().Add("Content-Type", "application/json")
			resp, _ := json.Marshal(data)
			w.Write(resp)
		}
	}
}

func PostWrapHandler(handler HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		bodyData, err := handler(w, req)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		} else {
			w.Header().Add("Content-Type", "application/json")
			resp, _ := json.Marshal(bodyData)
			w.Write(resp)
		}
	}
}

