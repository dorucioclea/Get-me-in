package internal

import (
	. "fmt"
	"github.com/ProjectReferral/Get-me-in/marketing-service/configs"
	"github.com/ProjectReferral/Get-me-in/marketing-service/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"io"
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

	dynamodb.DeleteItem(w, AdvertId(w, r.Body))
}

func GetAdvert(w http.ResponseWriter, r *http.Request) {

	dynamodb.GetItem(w, AdvertId(w, r.Body))
}

func UpdateAdvert(w http.ResponseWriter, r *http.Request) {

	//TODO: Change to UpdateItem
	dynamodb.CreateItem(w, 	dynamodb.DecodeToDynamoAttribute(w, r, models.Advert{}))
}

func AdvertId(w http.ResponseWriter, r io.ReadCloser) string{
	bodyMap := dynamodb.DecodeToMap(w, r, models.Advert{})
	return StringFromMap(bodyMap, configs.FIND_BY_ID)
}

func StringFromMap(m map[string]interface{}, p string) string{
	return Sprintf("%v", m[p])
}
