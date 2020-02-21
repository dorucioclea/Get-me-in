package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
* Item created using AttributeValue which is decoded by modeldecoding
*/
func CreateItem(av map[string]*dynamodb.AttributeValue) error{

	// translate into a compatible object
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(DynamoTable),
	}

	_, errM := DynamoConnection.PutItem(input)

	if errM != nil {
		return errM
	}

	return nil
}