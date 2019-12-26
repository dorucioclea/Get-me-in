package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Working")
	w.WriteHeader(http.StatusOK)
}

// var sess = session.Must(session.NewSessionWithOptions(session.Options{
// 	SharedConfigState: session.SharedConfigEnable,
// }))

type Item struct {
	Uuid      string `json:"id"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	//credentials.NewStaticCredentials("asd", "asd", "asd")
	var c = credentials.NewSharedCredentials("", "default")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-2"),
		Credentials: c,
	})

	var svc = dynamodb.New(sess)

	var u Item

	// Try to decode th
	//e request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	er1r := json.NewDecoder(r.Body).Decode(&u)

	if er1r != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//password := []byte(u.Password)

	//hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error hashing password")
	}

	// Create User using body info being passed in when the endpoint is called
	account := Item{
		Uuid:      u.Uuid,
		Firstname: u.Firstname,
		Surname:   u.Surname,
		Email:     u.Email,
		Password:  u.Password,
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
		Item:      av,
		TableName: aws.String("Users"),
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

// func DeleteUser(w http.ResponseWriter, r *http.Request) {

// 	var u internal.User

// 	// Try to decode the request body into the struct. If there is an error,
// 	// respond to the client with the error message and a 400 status code.
// 	err := json.NewDecoder(r.Body).Decode(&u)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	user_id := u.uuid

// 	input := &dynamodb.DeleteItemInput{
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"uuid": {
// 				N: aws.String(movieYear),
// 			},
// 		},
// 		TableName: aws.String("Users"),
// 	}

// 	_, err := svc.DeleteItem(input)

// 	if err != nil {
// 		w.WriteHeader(http.StatusFailedDependency)
// 		w.Write([]byte("424 - DynamoDB DeleteItem Failed"))
// 		return
// 	}

// 	fmt.Println("Deleted User from table.")
// }

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("")
// }

// func GetUser(w http.ResponseWriter, r *http.Request) {

// 	var u internal.User

// 	// Try to decode the request body into the struct. If there is an error,
// 	// respond to the client with the error message and a 400 status code.
// 	err := json.NewDecoder(r.Body).Decode(&u)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	result, err := svc.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String("Users"),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"uuid": {
// 				N: aws.String(u.uuid),
// 			},
// 		},
// 	})

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	item := Item{}
// 	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
// 	}

// 	if item.Title == "" {
// 		w.WriteHeader(http.StatusNoContent)
// 		w.Write([]byte("204 - Not user found"))
// 		fmt.Println("Could not find user")
// 		return
// 	}

// 	fmt.Println("User found: ", item)
// 	w.WriteHeader(http.StatusOK)
// 	return item
// }

// func Login(w http.ResponseWriter, r *http.Request) {
// 	var u internal.User

// 	// Try to decode the request body into the struct. If there is an error,
// 	// respond to the client with the error message and a 400 status code.
// 	err := json.NewDecoder(r.Body).Decode(&u)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	result, err := svc.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String("Users"),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"email": {
// 				N: aws.String(u.email),
// 			},
// 		},
// 	})

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	item := Item{}
// 	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
// 	}

// 	if item.Title == "" {
// 		w.WriteHeader(http.StatusNoContent)
// 		w.Write([]byte("204 - Not user found"))
// 		fmt.Println("Could not find user")
// 		return
// 	}

// 	loginCheck = bcrypt.CompareHashAndPassword(item.password, u.password)

// 	if loginCheck != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 	}
// 	w.WriteHeader(http.StatusAccepted)
// }

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("")
// }
