package main

import (
	"github.com/ProjectReferral/Get-me-in/account-api/configs"
	"github.com/ProjectReferral/Get-me-in/account-api/internal"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/api"
	"github.com/ProjectReferral/Get-me-in/account-api/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"log"
	"os"
)

func main() {
	loadEnvConfigs()

	internal.ConnectToDynamoDB()
	api.SetupEndpoints()
}

func loadEnvConfigs() {

	var env = ""

	log.Println("Running on %s \n", configs.PORT)

	configs.BrokerUrl = os.Getenv("broker_url")
	dynamodb.SearchParam = configs.UNIQUE_IDENTIFIER
	dynamodb.GenericModel = models.User{}

	switch env := os.Getenv("ENV"); env {
	case "DEV":
		dynamodb.DynamoTable = "dev-users"
	case "UAT":
		dynamodb.DynamoTable = "uat-users"
	case "PROD":
		dynamodb.DynamoTable = "prod-users"
	default:
		dynamodb.DynamoTable = "dev-users"
	}

	log.Println("Environment:" + env)
}

