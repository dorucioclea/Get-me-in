package internal

import (
	"encoding/json"
	"fmt"
	"github.com/ProjectReferral/Get-me-in/account-service/configs"
	"github.com/ProjectReferral/Get-me-in/account-service/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"net/http"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ConnectToInstance(w http.ResponseWriter, r *http.Request) {
	c := credentials.NewSharedCredentials("", "default")

	err := dynamodb.Connect(c, configs.EU_WEST_2)

	if err != nil {
		e := err.(*dynamodb.ErrorString)
		http.Error(w, e.Reason, e.Code)
	}

	w.WriteHeader(http.StatusOK)
}

// Tested & Working
func CreateUser(w http.ResponseWriter, r *http.Request) {
	dynamoAttr, errDecode := dynamodb.DecodeToDynamoAttribute(r.Body, models.User{})

	if !HandleError(errDecode, w, false){

		err := dynamodb.CreateItem(dynamoAttr)

		if !HandleError(err, w, false){
			w.WriteHeader(http.StatusOK)
		}
	}
}

// Tested & Working
func GetUser(w http.ResponseWriter, r *http.Request) {
	result, err := dynamodb.GetItem(ExtractValue(w, r))

	if !HandleError(err, w, true) {
		b, err := json.Marshal(dynamodb.Unmarshal(result, models.User{}))

		if !HandleError(err, w, false){

			w.Write([]byte(b))
			w.WriteHeader(http.StatusOK)
		}
	}
}

// Tested & Working
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	extractValue := ExtractValue(w, r)

	errDelete := dynamodb.DeleteItem(extractValue)

	if !HandleError(errDelete, w, false) {

		//Check item still exists
		result, err := dynamodb.GetItem(extractValue)

		//error thrown, record not found
		if !HandleError(err, w, true) {
			http.Error(w, result.GoString(), 302)
		}
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	//TODO: Change to UpdateItem
	CreateUser(w,r)}

//TODO: Refactor
func Login(w http.ResponseWriter, r *http.Request) {
	//
	//bodyMap,err := dynamodb.DecodeToMap(r.Body, models.Credentials{})
	//
	//if err != nil{
	//	return nil, err
	//}
	//
	//av, errM := dynamodbattribute.MarshalMap(bodyMap)
	//
	//
	//
	//
	//
	//bodyEmail := StringFromMap(bodyMap, configs.QUERY_PARAM)
	//bodyPassword := StringFromMap(bodyMap, configs.PW)
	//
	//result, found := dynamodb.GetItem(w, bodyEmail)
	//
	//u := dynamodb.Unmarshal(result, models.Credentials{})
	//dbPassword := StringFromMap(u, configs.PW)
	//
	//if found {
	//	if bodyPassword == dbPassword {
	//		w.WriteHeader(http.StatusAccepted)
	//	}
	//	w.WriteHeader(http.StatusUnauthorized)
	//}
	//w.WriteHeader(http.StatusNoContent)
}

func StringFromMap(m map[string]interface{}, p string) string{
	return fmt.Sprintf("%v", m[p])
}

func ExtractValue(w http.ResponseWriter, r *http.Request) string{

	v, err := dynamodb.GetParameterValue(r.Body, models.User{})
	HandleError(err, w, false)

	return v
}
