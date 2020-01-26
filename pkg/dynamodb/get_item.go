package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

func GetItem(itemValue string) (*dynamodb.GetItemOutput, error) {

	result, err := DynamoConnection.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(DynamoTable),
		Key: map[string]*dynamodb.AttributeValue{
			SearchParam: {
				S: aws.String(itemValue),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, &ErrorString{
			Reason: http.StatusText(http.StatusNotFound),
			Code:   http.StatusNotFound,
		}
	}

	return result, nil
}