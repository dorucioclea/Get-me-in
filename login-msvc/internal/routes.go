package internal

import (
	"net/http"
)

func SetupEndpoints(){

	http.HandleFunc("/auth", VerifyCredentials)
	http.HandleFunc("/mock", MockResponse)
}