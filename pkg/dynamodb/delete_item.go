package dynamodb

import (
	"net/http"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteItem(w http.ResponseWriter, identifier string) bool {

	fmt.Println(SearchParam)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
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
