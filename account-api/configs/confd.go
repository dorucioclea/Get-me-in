package configs

const (
	PORT = ":5002"

	/************** DynamoDB configs *************/
	EU_WEST_2 = "eu-west-2"
	UNIQUE_IDENTIFIER = "email"
	PW = "password"
	/*********************************************/
	/************** RabbitMQ configs *************/
	FANOUT_EXCHANGE = "accounts.fanout"
	/*********************************************/
	/*********** Authentication configs **********/
	AUTH_REGISTER = "register_user"
	AUTH_AUTHENTICATED = "read_update_delete"
	AUTH_LOGIN = "sigin_user"
	/*********************************************/
)

var (
	BrokerUrl = ""
)