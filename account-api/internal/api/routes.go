package api

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func wrapHandlerWithSpecialAuth(handler http.HandlerFunc, claim string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		if a != "" && security.VerifyTokenWithClaim(a, claim) {
			handler(w,req)
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}

func SetupEndpoints() {

	_router := mux.NewRouter()

	_router.HandleFunc("/test", TestFunc)

	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(CreateUser, configs.AUTH_REGISTER)).Methods("PUT")
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(DeleteUser, configs.AUTH_AUTHENTICATED)).Methods("DELETE")
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(UpdateUser, configs.AUTH_AUTHENTICATED)).Methods("PATCH")
	_router.HandleFunc("/account", wrapHandlerWithSpecialAuth(GetUser, configs.AUTH_AUTHENTICATED)).Methods("GET")
	//_router.HandleFunc("/account", GetUsers).Methods("GET")
	_router.HandleFunc("/account/verify", wrapHandlerWithSpecialAuth(Login, configs.AUTH_LOGIN)).Methods("POST")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}