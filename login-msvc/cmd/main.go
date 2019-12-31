package main

import (
	"../configs"
	"../internal"
	"net/http"
)

func main() {
	internal.SetupEndpoints()
	http.ListenAndServe(configs.PORT, nil)
}