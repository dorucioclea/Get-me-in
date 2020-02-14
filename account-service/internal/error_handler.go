package internal

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