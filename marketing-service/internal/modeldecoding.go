package internal

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github/Get-me-in/marketing-service/internal/models"
	"io"
	"net/http"
)

func DecodeToDynamoAttribute(w http.ResponseWriter, r *http.Request) map[string]*dynamodb.AttributeValue{

	av, errM := dynamodbattribute.MarshalMap(DecodeToJSON(w, r.Body))

	if errM != nil {
		http.Error(w, errM.Error(), http.StatusFailedDependency)
		w.Write([]byte("424 - DynamoDB Marshalling Failed"))
	}

	return av
}

func DecodeToJSON(w http.ResponseWriter, b io.ReadCloser ) models.Advert{

	var u models.Advert
	// Try to decode th
	//e request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	errJson := json.NewDecoder(b).Decode(&u)

	if errJson != nil {
		http.Error(w, errJson.Error(), http.StatusBadRequest)
	}

	return u
}