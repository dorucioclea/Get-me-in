package main

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/account-service/configs"
	"github.com/ProjectReferral/Get-me-in/account-service/internal/models"
	q_helper "github.com/ProjectReferral/Get-me-in/account-service/internal/q-helper"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"os"
)

func main() {
	loadEnvConfigs()
	//internal.SetupEndpoints()


	q_helper.ReceiveFromQ()

}

//TODO: improve workflow
func loadEnvConfigs() {

	var env = ""

	fmt.Printf("Running on %s \n", configs.PORT)

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

	fmt.Println("Environment:" + env)

}