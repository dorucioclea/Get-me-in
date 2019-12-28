package dynamodb

import "net/http"

/*
import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	var u internal.User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {
				N: aws.String(u.SearchParam),
			},
		},
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	item := Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if item.Title == "" {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("204 - Not user found"))
		fmt.Println("Could not find user")
		return
	}

	fmt.Println("User found: ", item)
	w.WriteHeader(http.StatusOK)
	return item
}
*/
func GetItem(w http.ResponseWriter, s string, v string) (string, bool){
	return "userdetails",true
}