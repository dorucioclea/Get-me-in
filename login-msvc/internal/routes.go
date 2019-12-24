package internal

import (
	"net/http"
)

func SetupEndpoints(){

	http.HandleFunc("/", Connection)
	http.HandleFunc("/headers", Headers)
	http.HandleFunc("/get", GET)
}