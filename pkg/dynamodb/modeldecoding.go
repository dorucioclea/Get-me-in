package dynamodb

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"io"
	"net/http"
)

func DecodeToDynamoAttribute(w http.ResponseWriter, r *http.Request, m interface{}) map[string]*dynamodb.AttributeValue{

	av, errM := dynamodbattribute.MarshalMap(DecodeToMap(w, r.Body, m))

	if errM != nil {
		http.Error(w, errM.Error(), http.StatusFailedDependency)
		w.Write([]byte("424 - DynamoDB Marshalling Failed"))
	}

	return av
}

func DecodeToMap (w http.ResponseWriter, b io.ReadCloser, m interface{}) map[string]interface{} {

	// Try to decode th
	//e request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	errJson := json.NewDecoder(b).Decode(&m)

	if errJson != nil {
		http.Error(w, errJson.Error(), http.StatusBadRequest)
	}

	mapM, ok := m.(map[string]interface{})

	if !ok {
		fmt.Printf("ERROR: not a map-> %#v\n", m)
	}

	return mapM
}

func Unmarshal(result *dynamodb.GetItemOutput, m interface{}) map[string]interface{} {

	err := dynamodbattribute.UnmarshalMap(result.Item, &m)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	mapM, ok := m.(map[string]interface{})

	if !ok {
		fmt.Printf("ERROR: not a map-> %#v\n", m)
	}

	return mapM
}