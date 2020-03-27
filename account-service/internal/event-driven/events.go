package event_driven

import "github.com/ProjectReferral/Get-me-in/account-service/configs"

func BroadcastUserCreatedEvent(body string){
	uId := NewUUID()

	//send to fanout exchange
	SendToQ(configs.ROUTING_KEY_RPOSTUSER, body, configs.EXCHANGE, uId)
}
