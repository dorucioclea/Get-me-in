package event_driven

import (
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter, isCustom bool) bool{

	if err != nil {
		if isCustom {
			e := err.(*dynamodb.ErrorString)
			http.Error(w, e.Reason, e.Code)
			return true
		}
		http.Error(w, err.Error(), 400)
		return true
	}
	return false
}

func HandleErrorEvent(err error, correlationId string, routingKey string, isCustom bool) bool{

	if err != nil {
		if isCustom {
			e := err.(*dynamodb.ErrorString)
			SendToQ(routingKey, string(e.Code) + e.Reason,  "account", correlationId)
			return true
		}

		SendToQ(routingKey, err.Error(),  "account", correlationId)
		return true
	}
	return false
}