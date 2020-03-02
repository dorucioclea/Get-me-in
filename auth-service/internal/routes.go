package internal

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/auth-service/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func
SetupEndpoints(){
	_router := mux.NewRouter()

	_router.HandleFunc("/auth", wrapHandlerWithAuth(VerifyCredentials)).Methods("GET")
	_router.HandleFunc("/mock", wrapHandlerWithAuth(MockResponse)).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}

func wrapHandlerWithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		fmt.Println(a,"asd")
		if a != "" && security.VerifyToken(a) {
			handler(w,req)
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}