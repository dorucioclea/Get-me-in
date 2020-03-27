package configs

const (
	PORT = ":5002"

	/*** DynamoDB configs ****/
	EU_WEST_2 = "eu-west-2"
	UNIQUE_IDENTIFIER = "email"
	PW = "password"
	/*************************/

	/**** RabbitMQ configs *****/
	EXCHANGE = "account"

	Q_POSTUSER = "create_user"
	Q_GETUSER = "read_user"
	Q_DELETEUSER = "delete_user"

	ROUTING_KEY_RPOSTUSER = "user.create.reply"
	ROUTING_KEY_RGETUSER = "user.read.reply"
	ROUTING_KEY_RDELETEUSER = "user.delete.reply"
	/*********************************************/

)

var (
	BrokerUrl = ""
)