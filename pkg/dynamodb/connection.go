package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DynamoTable string
var SearchParam string
var DynamoConnection *dynamodb.DynamoDB
var GenericModel interface{}

func Connect(c *credentials.Credentials, r string) error {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(r),
		Credentials: c,
	})

	if err != nil {
		return err
	}

	DynamoConnection = dynamodb.New(sess)

	return nil
}
