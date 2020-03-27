package main

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/pkg/rabbit-mq"
	"github.com/ProjectReferral/Get-me-in/pkg/rabbit-mq/configs"
	"os"
)

func main() {
	fmt.Println("Rabbit-Mq Handler")
	configs.BrokerUrl = os.Getenv("broker_url")

	fmt.Println(configs.BrokerUrl)

	rabbit_mq.SendTest("test", "custom message from GOlang", configs.TESTQ, "test.direct")


	forever := make(chan string)

	rabbit_mq.ReceiveFromQ(configs.TESTQ)
	rabbit_mq.ReceiveFromQ(configs.TESTQ1)


	<-forever
}