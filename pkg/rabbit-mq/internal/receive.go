package internal

import (
	"github.com/ProjectReferral/Get-me-in/pkg/rabbit-mq/configs"
	"github.com/streadway/amqp"
	"log"
)

func ReceiveFromQ(qName string){
	conn, err := amqp.Dial(configs.BrokerUrl)
	log.Printf("Listening on Q: %s", qName)

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		qName, // queue
		"",     // consumer
		true,   // auto-ack, TODO: manual ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan string)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-forever

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}