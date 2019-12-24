package configs

type User struct {
	uuid      string `json:"id"`
	firstname string `json:"firstname"`
	surname   string `json:"surname"`
	email     string `json:"email"`
	password  string `json:"password"`
}
