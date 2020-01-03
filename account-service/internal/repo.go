package internal

import (
	"../configs"
	"pkg/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"l"
	"net/http"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ConnectToInstance(w http.ResponseWriter, r *http.Request) {

	c := credentials.NewSharedCredentials("", "default")

	dynamodb.Connect(w, c, configs.EU_WEST_2)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	dynamodb.CreateItem(w, DecodeToDynamoAttribute(w, r))
}

func Login(w http.ResponseWriter, r *http.Request){

	body, found := dynamodb.GetItem(w, configs.FIND_BY_EMAIL, DecodeToJSON(w, r.Body).Email)

	if found {
		w.Write([]byte(body))
	}
}