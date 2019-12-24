package main

import (
	"fmt"
	internal "github/Get-me-in/account-service/internal"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Working")

	_router := mux.NewRouter()

	_router.HandleFunc("/test", internal.TestFunc)

	// // Register
	_router.HandleFunc("/account", CreateUser).Methods("PUT")
	_router.HandleFunc("/account/{id}", DeleteUser).Methods("DELETE")
	_router.HandleFunc("/account", UpdateUser).Methods("PATCH")
	_router.HandleFunc("/account/{id}", GetUser).Methods("GET")
	_router.HandleFunc("/account", GetUsers).Methods("GET")
	_router.HandleFunc("/account/verify", Login).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8000", _router))

}
