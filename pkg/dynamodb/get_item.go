package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetItem(identifier string) (*dynamodb.GetItemOutput, error) {

	result, err := DynamoConnection.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(DynamoTable),
		Key: map[string]*dynamodb.AttributeValue{
			SearchParam: {
				S: aws.String(identifier),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}