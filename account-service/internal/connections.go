package internal

import (
	"github.com/ProjectReferral/Get-me-in/account-service/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection

func ConnectToDynamoDB(){

	c := credentials.NewSharedCredentials("", "default")

	err := dynamodb.Connect(c, configs.EU_WEST_2)

	if err != nil {
		panic(err)
	}
}
