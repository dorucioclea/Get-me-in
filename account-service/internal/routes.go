package internal

import (
	"github.com/gorilla/mux"
	"../configs"
	"log"
	"net/http"
)

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/test", TestFunc)

	// connect to the db
	_router.HandleFunc("/connect", ConnectToInstance)

	// // Register
	_router.HandleFunc("/account", CreateUser).Methods("PUT")
	// _router.HandleFunc("/account/{id}", DeleteUser).Methods("DELETE")
	// _router.HandleFunc("/account", UpdateUser).Methods("PATCH")
	// _router.HandleFunc("/account/{id}", GetUser).Methods("GET")
	// _router.HandleFunc("/account", GetUsers).Methods("GET")
	_router.HandleFunc("/account/verify", Login).Methods("POST")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}