package internal

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/marketing-service/configs"
	"github.com/ProjectReferral/Get-me-in/marketing-service/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"net/http"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
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

func CreateAdvert(w http.ResponseWriter, r *http.Request) {

	dynamoAttr, errDecode := dynamodb.DecodeToDynamoAttribute(r.Body, models.Advert{})

	if !HandleError(errDecode, w, false){

		err := dynamodb.CreateItem(dynamoAttr)

		if !HandleError(err, w, false){
			w.WriteHeader(http.StatusOK)
		}
	}
}

func DeleteAdvert(w http.ResponseWriter, r *http.Request) {

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

func GetAdvert(w http.ResponseWriter, r *http.Request) {

	result, err := dynamodb.GetItem(ExtractValue(w, r))

	if !HandleError(err, w, true) {
		b, err := json.Marshal(dynamodb.Unmarshal(result, models.Advert{}))

		if !HandleError(err, w, false){

			w.Write([]byte(b))
			w.WriteHeader(http.StatusOK)
		}
	}
}

func UpdateAdvert(w http.ResponseWriter, r *http.Request) {

	//TODO: Change to UpdateItem
	CreateAdvert(w,r)
}

func ExtractValue(w http.ResponseWriter, r *http.Request) string{

	v, err := dynamodb.GetParameterValue(r.Body, models.Advert{})
	HandleError(err, w, false)

	return v
}
