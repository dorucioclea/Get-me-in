package main

import (
	"fmt"
	"github.com/ProjectReferral/Get-me-in/pkg/rabbit-mq/configs"
	"github.com/ProjectReferral/Get-me-in/pkg/rabbit-mq/internal"
	"os"
)

func main() {
	fmt.Println("Rabbit-Mq Handler")
	configs.BrokerUrl = os.Getenv("broker_url")

	fmt.Println(configs.BrokerUrl)

	internal.SendToQ("test", "custom message from GOlang", configs.TESTQ, "test.direct")


	forever := make(chan string)

	internal.ReceiveFromQ(configs.TESTQ)
	internal.ReceiveFromQ(configs.TESTQ1)


	<-forever
}