package internal

import (
	"github.com/ProjectReferral/Get-me-in/pkg/security"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupEndpoints() *mux.Router{
	_router := mux.NewRouter()

	_router.HandleFunc("/test", wrapHandlerWithAuth(TestFunc))
	// connect to the db
	_router.HandleFunc("/connect", wrapHandlerWithAuth(ConnectToInstance))
	_router.HandleFunc("/advert", wrapHandlerWithAuth(CreateAdvert)).Methods("PUT")
	_router.HandleFunc("/advert", wrapHandlerWithAuth(DeleteAdvert)).Methods("DELETE")
	_router.HandleFunc("/advert", wrapHandlerWithAuth(UpdateAdvert)).Methods("PATCH")
	_router.HandleFunc("/advert", wrapHandlerWithAuth(GetAdvert)).Methods("GET")

	return _router
}

func wrapHandlerWithAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		a := req.Header.Get("Authorization")

		if a != "" && security.VerifyToken(a)  {
			handler(w,req)
		}

		w.WriteHeader(http.StatusUnauthorized)
	}
}


