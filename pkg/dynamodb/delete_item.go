package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteItem(itemValue string) error {

	// translate into a compatible object
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			SearchParam: {
				S: aws.String(itemValue),
			},
		},
		TableName: aws.String(DynamoTable),
	}

	_, err := DynamoConnection.DeleteItem(input)

	if err != nil {
		return err
	}

	return nil
}
