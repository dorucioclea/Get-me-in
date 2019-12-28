package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

func CreateItem(w http.ResponseWriter, av map[string]*dynamodb.AttributeValue) {

	// Adding item to database..
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(DynamoTable),
	}

	_, errM := DynamoConnection.PutItem(input)

	if errM != nil {
		http.Error(w, errM.Error(), http.StatusFailedDependency)
		w.Write([]byte("424 - DynamoDB PuTItem Failed"))
	}

	w.WriteHeader(http.StatusCreated)
}