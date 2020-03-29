package event_driven

/*func HandleErrorEvent(err error, correlationId string, routingKey string, isCustom bool) bool{

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
}*/