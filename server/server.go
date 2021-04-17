package main

import "github.com/gorilla/mux"

type server struct {
	router *mux.Router
}

func newServer() *server {
	router := mux.NewRouter()

	return &server{
		router: router,
	}
}
