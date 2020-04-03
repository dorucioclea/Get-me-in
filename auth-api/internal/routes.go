package internal

import (
	"github.com/ProjectReferral/Get-me-in/auth-api/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetupEndpoints(){
	_router := mux.NewRouter()

	_router.HandleFunc("/auth", VerifyCredentials).Methods("GET")
	_router.HandleFunc("/auth/temp", IssueRegistrationTempToken).Methods("GET")
	_router.HandleFunc("/mock", MockResponse).Methods("GET")

	log.Fatal(http.ListenAndServe(configs.PORT, _router))
}


//Not needed at the current stage, endpoints open to everyone
//Credentials validator and recaptcha already in place on endpoint
func wrapHandlerWithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		if a != "" && security.VerifyToken(a) {
			handler(w,req)
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}