package internal

import (
	"net/http"
)

func SetupEndpoints(){

	http.HandleFunc("/", APIInfo)
	http.HandleFunc("/headers", Headers)
	http.HandleFunc("/get", GET)
}