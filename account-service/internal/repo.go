package internal

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github/Get-me-in/account-service/configs"
	"github/Get-me-in/account-service/internal/models"
	"github/Get-me-in/pkg/dynamodb"
	"io"
	"log"
	"net/http"
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
	dynamodb.CreateItem(w, dynamodb.DecodeToDynamoAttribute(w, r, models.User{}))

}

// Tested & Working
func GetUser(w http.ResponseWriter, r *http.Request) {
	result, status := dynamodb.GetItem(w, Email(w, r.Body))
	if status {
		w.WriteHeader(http.StatusOK)

		b, err := json.Marshal(dynamodb.Unmarshal(result, models.User{}))
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(b))
	}
}

// Tested & Working
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	status := dynamodb.DeleteItem(w, Email(w, r.Body))
	if status {
		w.WriteHeader(http.StatusOK)
	}
}

// Temporary
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	dynamodb.CreateItem(w, dynamodb.DecodeToDynamoAttribute(w, r, models.User{}))
}

// Tested & Working
func Login(w http.ResponseWriter, r *http.Request) {

	bodyMap := dynamodb.DecodeToMap(w, r.Body, models.Credentials{})
	bodyEmail := StringFromMap(bodyMap, configs.QUERY_PARAM)
	bodyPassword := StringFromMap(bodyMap, configs.PW)

	result, found := dynamodb.GetItem(w, bodyEmail)

	u := dynamodb.Unmarshal(result, models.Credentials{})
	dbPassword := StringFromMap(u, configs.PW)

	if found {
		if bodyPassword == dbPassword {
			w.WriteHeader(http.StatusAccepted)
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusNoContent)
}

func Email(w http.ResponseWriter, r io.ReadCloser) string{
	bodyMap := dynamodb.DecodeToMap(w, r, models.Credentials{})
	return StringFromMap(bodyMap, configs.QUERY_PARAM)
}

func StringFromMap(m map[string]interface{}, p string) string{
	return fmt.Sprintf("%v", m[p])
}