package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

func DeleteItem(w http.ResponseWriter, identifier string) bool {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			SearchParam: {
				S: aws.String(identifier),
			},
		},
		TableName: aws.String(DynamoTable),
	}

	_, err := DynamoConnection.DeleteItem(input)

	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte("424 - DynamoDB DeleteItem Failed"))
		return false
	}

	w.WriteHeader(http.StatusOK)
	return true
}
