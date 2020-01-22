package dynamodb

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"io"
	"net/http"
)

func DecodeToDynamoAttribute(r *http.Request, m interface{}) (map[string]*dynamodb.AttributeValue, error){

	bodyMap, err := DecodeToMap(r.Body, m)

	if err != nil{
		return nil, err
	}

	av, errM := dynamodbattribute.MarshalMap(bodyMap)

	if errM != nil {
		return nil, errM
	}

	return av, nil
}

func DecodeToMap (b io.ReadCloser, m interface{}) (map[string]interface{}, error) {

	// Try to decode th
	//e request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	errJson := json.NewDecoder(b).Decode(&m)

	if errJson != nil {
		return nil, errJson
	}

	mapM, ok := m.(map[string]interface{})

	if !ok {
		fmt.Printf("ERROR: not a map-> %#v\n", m)
	}

	return mapM, nil
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

func GetParameterValue(r io.ReadCloser, m interface{}) (string,error){
	bodyMap, err := DecodeToMap(r, m)

	if err != nil{
		return "", err
	}

	return StringFromMap(bodyMap, SearchParam), nil
}

func StringFromMap(m map[string]interface{}, p string) string{
	return fmt.Sprintf("%v", m[p])
}