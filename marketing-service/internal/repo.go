package internal

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github/Get-me-in/marketing-service/configs"
	"github/Get-me-in/pkg/dynamodb"
	"net/http"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func ConnectToInstance(w http.ResponseWriter, r *http.Request) {

	c := credentials.NewSharedCredentials("", "default")

	dynamodb.Connect(w, c, configs.EU_WEST_2)
}

func CreateAdvert(w http.ResponseWriter, r *http.Request) {

	dynamodb.CreateItem(w, DecodeToDynamoAttribute(w, r))
}