package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Need to do
func UpdateItem(updatingField string, identifier string, updateVal string) (bool, error) {

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
		return false, err
	}
	
	return true, nil
}
