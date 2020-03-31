package models

type User struct {
	Uuid      string `json:"id"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Credentials struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}