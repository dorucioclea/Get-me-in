package rabbit_mq

import (
	"fmt"
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

	msgsClone, err := ch.Consume(
		configs.TESTQ1, // queue
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
			log.Printf("Received a message: %s - %s", d.Body,d.CorrelationId)

			//ProcessMessage(1)
			SendToQ("test.reply", "Reply: processed", "queue.test.reply", "test.direct")
		}
	}()

	go func() {
		for d := range msgsClone {
			log.Printf("Received a message: %s - %s", d.Body,d.CorrelationId)
			//	ProcessMessage(2)
			SendToQ("test.reply", "Reply: processed", "queue.test.reply", "test.direct")
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

func ProcessMessage(process int){

	switch process {
	case 1:
		fmt.Println("Process")
		break
	case 2:
		fmt.Println("Process-Two")
		break
	}

}