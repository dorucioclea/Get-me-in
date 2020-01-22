package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

var DynamoTable string
var SearchParam string
var DynamoConnection *dynamodb.DynamoDB
var GenericModel interface{}

func Connect(w http.ResponseWriter, c *credentials.Credentials, r string) {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(r),
		Credentials: c,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	DynamoConnection = dynamodb.New(sess)
}
