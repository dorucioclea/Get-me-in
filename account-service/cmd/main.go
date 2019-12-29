package main

import (
	"fmt"
	"github/Get-me-in/account-service/internal"
	"github/Get-me-in/pkg/dynamodb"
	"os"
)

func main() {
	loadEnvConfigs()
	internal.SetupEndpoints()
}

//TODO: improve workflow
func loadEnvConfigs() {

	fmt.Print("Running on ")

	dynamodb.SearchParam = "Email"

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