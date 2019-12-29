package main

import (
"github/Get-me-in/login-msvc/configs"
"github/Get-me-in/login-msvc/internal"
"net/http"
)

func main() {
	internal.SetupEndpoints()
	http.ListenAndServe(configs.PORT, nil)
}