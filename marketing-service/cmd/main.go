package main

import (
	"fmt"
	"../configs"
	"../internal"
	"../../pkg/dynamodb"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnvConfigs()
	log.Fatal(http.ListenAndServe(configs.PORT, internal.SetupEndpoints()))
}

//TODO: improve workflow
func loadEnvConfigs() {

	fmt.Print("Running on ")

	dynamodb.SearchParam = configs.FIND_BY_ACCOUNT

	switch env := os.Getenv("ENV"); env {
	case "DEV":
		dynamodb.DynamoTable = "dev-adverts"
		fmt.Println(env)
	case "UAT":
		dynamodb.DynamoTable = "uat-adverts"
		fmt.Println(env)
	case "PROD":
		dynamodb.DynamoTable = "prod-adverts"
		fmt.Println(env)

	default:
		dynamodb.DynamoTable = "dev-adverts"
		fmt.Println(env)
	}
}