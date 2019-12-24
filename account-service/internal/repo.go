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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	svc := dynamodb.New(sess)
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
	fmt.Println("")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
}
