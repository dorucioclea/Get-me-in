package internal

import (
	"github.com/ProjectReferral/Get-me-in/account-service/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func wrapHandlerWithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		if a != "" && security.VerifyToken(a) {
			handler(w,req)
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/test", TestFunc)
	_router.HandleFunc("/connect", wrapHandlerWithAuth(ConnectToInstance))

	// Register
	_router.HandleFunc("/account", wrapHandlerWithAuth(CreateUser)).Methods("PUT")
	_router.HandleFunc("/account", wrapHandlerWithAuth(DeleteUser)).Methods("DELETE")
	_router.HandleFunc("/account", wrapHandlerWithAuth(UpdateUser)).Methods("PATCH")
	_router.HandleFunc("/account", wrapHandlerWithAuth(GetUser)).Methods("GET")
	//_router.HandleFunc("/account", GetUsers).Methods("GET")
	_router.HandleFunc("/account/verify", wrapHandlerWithAuth(Login)).Methods("POST")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}