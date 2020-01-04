package main

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/account-service/internal"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"os"
)

func main() {
	loadEnvConfigs()
	internal.SetupEndpoints()
}

//TODO: improve workflow
func loadEnvConfigs() {

	fmt.Print("Running on ")

	dynamodb.SearchParam = "email"

	switch env := os.Getenv("ENV"); env {
	case "DEV":
		dynamodb.DynamoTable = "dev-users"
		fmt.Println(env)
	case "UAT":
		dynamodb.DynamoTable = "uat-users"
		fmt.Println(env)
	case "PROD":
		dynamodb.DynamoTable = "prod-users"
		fmt.Println(env)

	default:
		dynamodb.DynamoTable = "dev-users"
		fmt.Println(env)
	}
}