package internal

import (
	"encoding/json"
	"github.com/ProjectReferral/Get-me-in/marketing-service/configs"
	"github.com/ProjectReferral/Get-me-in/marketing-service/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"net/http"
)

func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func ConnectToInstance(w http.ResponseWriter, r *http.Request) {

	c := credentials.NewSharedCredentials("", "default")

	err := dynamodb.Connect(c, configs.EU_WEST_2)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}

func CreateAdvert(w http.ResponseWriter, r *http.Request) {

	dynamoAttr , errDecode := dynamodb.DecodeToDynamoAttribute(r.Body, models.Advert{})
	HandleError(errDecode, w)

	err := dynamodb.CreateItem(dynamoAttr)
	HandleError(err, w)

	w.WriteHeader(http.StatusOK)
}

func DeleteAdvert(w http.ResponseWriter, r *http.Request) {

	//TODO: handle items not existent in db
	errDelete := dynamodb.DeleteItem(ExtractValue(w, r))
	HandleError(errDelete, w)

	w.WriteHeader(http.StatusOK)
}

func GetAdvert(w http.ResponseWriter, r *http.Request) {

	//TODO: handle items not existent in db
	result, err := dynamodb.GetItem(ExtractValue(w, r))
	HandleError(err, w)

	b, err := json.Marshal(dynamodb.Unmarshal(result, models.Advert{}))
	HandleError(err, w)

	w.Write([]byte(b))
	w.WriteHeader(http.StatusOK)
}

func UpdateAdvert(w http.ResponseWriter, r *http.Request) {

	//TODO: Change to UpdateItem
	CreateAdvert(w,r)
}

func ExtractValue(w http.ResponseWriter, r *http.Request) string{

	v, err := dynamodb.GetParameterValue(r.Body, models.Advert{})
	HandleError(err, w)

	return v
}
