package internal

import (
	"github.com/gorilla/mux"
	"../../pkg/security"
	"net/http"
)

func SetupEndpoints() *mux.Router{
	_router := mux.NewRouter()

	_router.HandleFunc("/test", wrapHandlerWithAuth(TestFunc))
	// connect to the db
	_router.HandleFunc("/connect", wrapHandlerWithAuth(ConnectToInstance))
	_router.HandleFunc("/account", wrapHandlerWithAuth(CreateAdvert)).Methods("PUT")

	return _router
}

func wrapHandlerWithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		if security.VerifyToken(a) && a != "" {
			handler(w,req)
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}


