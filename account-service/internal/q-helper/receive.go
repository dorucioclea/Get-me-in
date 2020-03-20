package q_helper

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/account-service/configs"
	"github.com/streadway/amqp"
	"log"
)

func ReceiveFromQ(){
	conn, err := amqp.Dial(configs.BrokerUrl)
	log.Printf("Listening on Q: %s", configs.Q_POSTUSER)

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")

	msgsCreateUser, err := ch.Consume(
		configs.Q_POSTUSER, // queue
		"",     // consumer
		true,   // auto-ack, TODO: manual ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	msgsGetUser, err := ch.Consume(
		configs.Q_GETUSER, // queue
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
		for d := range msgsCreateUser {
			log.Printf("Received a message: %s - %s", d.Body,d.CorrelationId)

			//ProcessMessage(1)
			SendToQ("user.create.reply", "Reply: processed", "create_advert_reply", "account")
		}
	}()

	go func() {
		for d := range msgsGetUser {
			log.Printf("Received a message: %s - %s", d.Body,d.CorrelationId)
			//	ProcessMessage(2)
			SendToQ("user.read.reply", "Reply: processed", "read_advert_reply", "account")
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
