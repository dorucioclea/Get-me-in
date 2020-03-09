package internal

import (
	"github.com/ProjectReferral/Get-me-in/pkg/rabbit-mq/configs"
	"github.com/streadway/amqp"
)

func SendToQ(routingKey string, body string, qName string, exchange string){
	conn, err := amqp.Dial(configs.BrokerUrl)

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		exchange,     // exchange
		routingKey, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish a message")
}
