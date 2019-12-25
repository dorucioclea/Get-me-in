package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"

	internal "github/Get-me-in/account-service/configs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/nu7hatch/gouuid"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Working")
	w.WriteHeader(http.StatusOK)
}

var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

svc := dynamodb.New(sess)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	// Create User using body info being passed in when the endpoint is called
	account := internal.User{
		uuid:     uuid.NewV4(),
		fistname: body.firstname,
		surname:  body.surname,
		email:    body.email,
		password: body.password,
	}

	// Converting shit to dynamo type
	av, err := dynamodbattribute.MarshalMap(account)

	if err != nil {
		fmt.Println("Got error marshalling user..")
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte("424 - DynamoDB Marshalling Failed"))
	}

	// Adding user to database..
	input := &dynamodb.PutItemInput{
		User:      av,
		TableName: aws.String("GetUsers"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte("424 - DynamoDB PuTItem Failed"))
	}
	w.WriteHeader(http.StatusCreated)
}

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
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
}

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
            	N: aws.String(u.uuid),
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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
}
