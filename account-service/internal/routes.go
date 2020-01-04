package internal

import (
	"github.com/gorilla/mux"
	"../configs"
	"log"
	"net/http"
	"github/Get-me-in/pkg/security"
	"fmt"
)

func wrapHandlerWithaAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		fmt.Println(a,"asd")
		if a != "" && security.VerifyToken(a) {
			handler(w,req)
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/test", TestFunc)
	_router.HandleFunc("/connect", ConnectToInstance)

	// Register
	_router.HandleFunc("/account", CreateUser).Methods("PUT")
	_router.HandleFunc("/account", wrapHandlerWithaAuth(DeleteUser)).Methods("DELETE")
	_router.HandleFunc("/account", wrapHandlerWithaAuth(UpdateUser)).Methods("PATCH")
	_router.HandleFunc("/account", GetUser).Methods("GET")
	//_router.HandleFunc("/account", GetUsers).Methods("GET")
	_router.HandleFunc("/account/verify", Login).Methods("POST")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}