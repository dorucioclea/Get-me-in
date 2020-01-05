package main

import (
	"github.com/ProjectReferral/Get-me-in/login-msvc/configs"
	"github.com/ProjectReferral/Get-me-in/login-msvc/internal"
	"net/http"
)

func main() {
	internal.SetupEndpoints()
	http.ListenAndServe(configs.PORT, nil)
}