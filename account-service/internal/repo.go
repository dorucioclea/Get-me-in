package internal

import (
	"github/Get-me-in/account-service/configs"
	"github/Get-me-in/pkg/dynamodb"
	"net/http"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"fmt"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ConnectToInstance(w http.ResponseWriter, r *http.Request) {
	c := credentials.NewSharedCredentials("", "default")
	dynamodb.Connect(w, c, configs.EU_WEST_2)
}

// Tested & Working
func CreateUser(w http.ResponseWriter, r *http.Request) {
	dynamodb.CreateItem(w, DecodeToDynamoAttribute(w, r))
}

// Tested & Working
func GetUser(w http.ResponseWriter, r *http.Request) {
	result, status := dynamodb.GetItem(w, DecodeToJSON(w, r.Body).Email)
	if status {
		Unmarshall(result)
		w.WriteHeader(http.StatusOK)
		fmt.Println(Unmarshall(result))
	}
}

// Tested & Working
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	status := dynamodb.DeleteItem(w, DecodeToJSON(w, r.Body).Email)
	if status {
		w.WriteHeader(http.StatusOK)
	}
}

// Temporary
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	dynamodb.CreateItem(w, DecodeToDynamoAttribute(w, r))
}

// Tested & Working
func Login(w http.ResponseWriter, r *http.Request) {

	j := DecodeToJSON(w, r.Body)
	result, found := dynamodb.GetItem(w, j.Email) 

	if found {
		fmt.Println(Unmarshall(result).Password, j.Password)
		if Unmarshall(result).Password == j.Password {
			w.WriteHeader(http.StatusAccepted)
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusNoContent)
}
