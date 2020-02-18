package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

/*** Values injected from main service that imports this library ***/
var DynamoTable string
var SearchParam string
var GenericModel interface{}
/*******************************************************************/

var DynamoConnection *dynamodb.DynamoDB

/*
* Create a connection to DB and assign the session to DynamoConnection variable
* DynamoConnection variable is shared by other resources(CRUD)
*/
func Connect(c *credentials.Credentials, region string) error {

	if DynamoTable == "" && SearchParam == "" && GenericModel == nil{
		return &ErrorString{
			Reason: "Injected values are empty or nil",
			Code:   http.StatusBadRequest,
		}
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: c,
	})

	if err != nil {
		return err
	}

	DynamoConnection = dynamodb.New(sess)

	return nil
}
