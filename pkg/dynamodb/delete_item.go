package dynamodb

/*
import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var u internal.User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user_id := u.uuid

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {
				N: aws.String(movieYear),
			},
		},
		TableName: aws.String("Users"),
	}

	_, err := svc.DeleteItem(input)

	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte("424 - DynamoDB DeleteItem Failed"))
		return
	}

	fmt.Println("Deleted User from table.")
}*/