package models

type Advert struct {
	Uuid      	string `json:"id"`
	AccountId	string `json:"account_id"`
	Title   	string `json:"title"`
	MaxUsers 	string `json:"max_users"`
	Premium		bool   `json:"premium"`
	ValidFrom   string `json:"valid_from"`
	ValidTill	string `json:"valid_till"`
	Company  	string `json:"company"`
	Description string `json:"description"`
}

