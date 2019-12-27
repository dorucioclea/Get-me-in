package main

import (
	"fmt"
	"github/Get-me-in/account-service/configs"
	"github/Get-me-in/account-service/internal"
	"os"
)

func main() {
	loadEnvConfigs()
	internal.SetupEndpoints()
}

func loadEnvConfigs() {

	fmt.Print("Running on ")
	switch env := os.Getenv("ENV"); env {
	case "DEV":
		configs.DYNAMO_TABLE = "dev-users"
		fmt.Println(env)

	case "UAT":
		configs.DYNAMO_TABLE = "uat-users"
		fmt.Println(env)

	case "PROD":
		configs.DYNAMO_TABLE = "uat-users"
		fmt.Println(env)

	default:
		configs.DYNAMO_TABLE = "dev-users"
		fmt.Println(env)
	}
}