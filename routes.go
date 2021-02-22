package main

import (
	"github.com/gorilla/mux"
)

func setupRoutesConfig() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RunNetworkHandler).Methods("GET")
	return router
}
