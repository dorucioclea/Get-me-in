package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateItem(av map[string]*dynamodb.AttributeValue) error{

	// Adding item to database..
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