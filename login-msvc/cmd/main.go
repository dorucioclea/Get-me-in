package main

import (
	"Get-me-in/login-msvc/configs"
	"Get-me-in/login-msvc/internal"
	"net/http"
)

func main() {
	internal.SetupEndpoints()
	http.ListenAndServe(configs.PORT, nil)
}