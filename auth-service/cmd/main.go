package main

import (
	"github.com/ProjectReferral/Get-me-in/auth-service/configs"
	"github.com/ProjectReferral/Get-me-in/auth-service/internal"
	"net/http"
)

func main() {
	internal.SetupEndpoints()
	http.ListenAndServe(configs.PORT, nil)
}