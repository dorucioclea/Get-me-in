package event_driven

import "github.com/ProjectReferral/Get-me-in/account-service/configs"

func UserCreatedEvent(body string){
	uId := NewUUID()
	SendToQ(configs.ROUTING_KEY_RPOSTUSER, body, configs.EXCHANGE, uId)
}
