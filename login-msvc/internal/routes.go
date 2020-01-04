package internal

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"net/http"
)

func SetupEndpoints(){
	http.HandleFunc("/auth", VerifyCredentials)
	http.HandleFunc("/mock", MockResponse)
}

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