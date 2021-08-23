package server

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", GetUsers()).Methods("GET")
	router.HandleFunc("/employees", GetEmployees()).Methods("GET")

	return router
}
