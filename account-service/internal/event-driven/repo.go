package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/account-service/configs"
	"github.com/ProjectReferral/Get-me-in/account-service/internal/models"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
)

func CreateUser(userJson []byte, correlationId string) bool{

	dynamoAttr, errDecode := dynamodb.DecodeToDynamoAttributeFromByte(userJson, models.User{})

	if !HandleErrorEvent(errDecode, configs.ROUTING_KEY_RPOSTUSER, correlationId, false) {

		err := dynamodb.CreateItem(dynamoAttr)

		HandleErrorEvent(err, configs.ROUTING_KEY_RPOSTUSER, correlationId,  false)
	}

	return true
}

