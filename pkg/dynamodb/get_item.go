package dynamodb

import (
	"fmt"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetItem(w http.ResponseWriter, identifier string) (*dynamodb.GetItemOutput, bool) {
	fmt.Println(identifier, SearchParam, DynamoTable)
	result, err := DynamoConnection.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(DynamoTable),
		Key: map[string]*dynamodb.AttributeValue{
			SearchParam: {
				S: aws.String(identifier),
			},
		},
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	return result, true
}