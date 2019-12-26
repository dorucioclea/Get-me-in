package models

type Message struct {
	ApiVersion string
	Platform string
	Version string
	Token string
}

type Client struct {
	Username string
	Password string
}