package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

//Need to do
func UpdateItem(w http.ResponseWriter, updatingField string, identifier string, updateVal string) bool {

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			updatingField: {
				N: aws.String(updateVal),
			},
		},
		TableName: aws.String(DynamoTable),
		Key: map[string]*dynamodb.AttributeValue{
			SearchParam: {
				N: aws.String(identifier),
			},
		},
		// May need updating
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rating = :r"),
	}

	_, err := DynamoConnection.UpdateItem(input)
	if err != nil {
		return false
	}
	
	return true
}
