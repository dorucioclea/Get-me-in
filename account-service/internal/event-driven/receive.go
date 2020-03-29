package event_driven

import "log"

/*
func ReceiveFromAllQs(){
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
		false,   // auto-ack, TODO: manual ack
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
			//Debugging purposes
			log.Printf("Received a message: %s - %s", d.Body,d.CorrelationId)


			d.Reject(true)


			if CreateUser(d.Body, d.CorrelationId) {
				SendToQ(configs.ROUTING_KEY_RPOSTUSER, "Reply: User created for "+d.CorrelationId,  configs.EXCHANGE, d.CorrelationId)
			}
		}
	}()

	go func() {
		for d := range msgsGetUser {
			log.Printf("Received a message: %s - %s", d.Body,d.CorrelationId)
			//	ProcessMessage(2)
			SendToQ("user.read.reply", "Reply: processed",  "account", d.CorrelationId)
		}
	}()

	<-forever

	//Debugging purposes
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
*/

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

