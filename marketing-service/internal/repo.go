package internal

import (
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

	dynamodb.Connect(w, c, configs.EU_WEST_2)
}

func CreateAdvert(w http.ResponseWriter, r *http.Request) {

	dynamodb.CreateItem(w, dynamodb.DecodeToDynamoAttribute(w, r, models.Advert{}))
}

func DeleteAdvert(w http.ResponseWriter, r *http.Request) {

	dynamodb.DeleteItem(w, dynamodb.GetParameterValue(w, r.Body, models.Advert{}))
}

func GetAdvert(w http.ResponseWriter, r *http.Request) {

	dynamodb.GetItem(w, dynamodb.GetParameterValue(w, r.Body, models.Advert{}))
}

func UpdateAdvert(w http.ResponseWriter, r *http.Request) {

	//TODO: Change to UpdateItem
	dynamodb.CreateItem(w, 	dynamodb.DecodeToDynamoAttribute(w, r, models.Advert{}))
}


